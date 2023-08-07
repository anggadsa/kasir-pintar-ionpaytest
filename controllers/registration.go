package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kasir-pintar-ionpaytest/core"
	"kasir-pintar-ionpaytest/models"
	"kasir-pintar-ionpaytest/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Registration struct {
	core.Controller
}

var success Success

func (a *Registration) Index(c *gin.Context) {
	var params *models.Registration
	c.BindJSON(&params)
	result, list := new(services.Registration).All(params)
	a.Json(c, result, list)
}

func writeLog(nameReq string, body string, nameRes string) {
	if nameRes == "request" { // Request Log
		var params *models.Registration
		// fmt.Println("MASUK", body)
		json.Unmarshal([]byte(body), &params)
		runLogFile, _ := os.OpenFile(
			"./log/"+nameReq+"/"+nameRes+".log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
		log.Logger = zerolog.New(multi)
		log.Info().
			Timestamp().
			Interface("Data", params).Msg("")

	} else if nameRes == "response" { // Response Log
		var success Success
		// fmt.Println("MASUK", body)
		json.Unmarshal([]byte(body), &success)

		runLogFile, _ := os.OpenFile(
			"./log/"+nameReq+"/"+nameRes+".log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
		log.Logger = zerolog.New(multi)
		log.Info().
			Timestamp().
			Interface("Data", Success(success)).Msg("")
	} else if nameRes == "reqStatus" {
		var status Status
		json.Unmarshal([]byte(body), &status)
		runLogFile, _ := os.OpenFile(
			"./log/status/request.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
		log.Logger = zerolog.New(multi)
		log.Info().
			// Str("Data", body).
			Timestamp().
			Interface("Data", Status(status)).Msg("")
	} else if nameRes == "reqPayment" {
		var param Param
		json.Unmarshal([]byte(body), &param)
		runLogFile, _ := os.OpenFile(
			"./log/payment/request.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
		log.Logger = zerolog.New(multi)
		log.Info().
			// Str("Data", body).
			Timestamp().
			Interface("Data", Param(param)).Msg("")
	} else {
		runLogFile, _ := os.OpenFile(
			"./log/payment/response.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
		log.Logger = zerolog.New(multi)
		log.Info().
			Str("Data", body).
			Timestamp().
			Msg("")

	}
}

func (a *Registration) Store(c *gin.Context) {
	var params *models.Registration
	c.BindJSON(&params)

	// Save Response Log
	reqBytes, _ := json.Marshal(&params)
	fmt.Println(string(reqBytes))
	// fmt.Println(params)
	writeLog("registration", string(reqBytes), "request")

	// Save to Database
	result, _ := new(services.Registration).Create(params)

	// Initialization HTTP Request to Nice Pay API Registration
	decode, _ := json.Marshal(result.Value)
	url := "https://dev.nicepay.co.id/nicepay/direct/v2/registration"
	method := "POST"
	payload := bytes.NewReader(decode)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	// Add Header to HTTP Request
	req.Header.Add("Content-Type", "application/json")
	// Send HTTP Request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()

	// Read Body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Save Response Log
	writeLog("registration", string(body), "response")

	// Send Response to User
	json.Unmarshal([]byte(body), &success)
	a.Json(c, result, success)
}

// func (a *Registration) Show(c *gin.Context) {

// 	data, err := io.ReadAll(c.Request.Body)
// 	// fmt.Println(data)

// 	url := "https://dev.nicepay.co.id/nicepay/direct/v2/inquiry"
// 	method := "POST"
// 	payload := bytes.NewReader(data)
// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	req.Header.Add("Content-Type", "application/json")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(body))
// 	var success Success
// 	json.Unmarshal([]byte(body), &success)

// 	// Logging
// 	runLogFile, _ := os.OpenFile(
// 		"./log/inquiry.log",
// 		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
// 		0664,
// 	)
// 	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
// 	log.Logger = zerolog.New(multi)
// 	log.Info().
// 		Timestamp().
// 		Interface(success.ResultMsg, Success(success)).Msg("")

// 	c.JSON(http.StatusOK, success)
// }
