package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kasir-pintar-ionpaytest/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Inquiry struct {
	core.Controller
}

type Status struct {
	TimeStamp     string `json:"timeStamp"`
	TXid          string `json:"tXid"`
	IMid          string `json:"iMid"`
	ReferenceNo   string `json:"referenceNo"`
	Amt           string `json:"amt"`
	MerchantToken string `json:"merchantToken"`
}

func (a *Inquiry) Show(c *gin.Context) {
	var status Status
	c.BindJSON(&status)
	reqBytes, _ := json.Marshal(&status)

	writeLog("status", string(reqBytes), "reqStatus")

	url := "https://dev.nicepay.co.id/nicepay/direct/v2/inquiry"
	method := "POST"
	payload := bytes.NewReader(reqBytes)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var success Success
	json.Unmarshal([]byte(body), &success)

	// Save Response Log
	writeLog("status", string(body), "response")

	c.JSON(http.StatusOK, success)
}
