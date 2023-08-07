package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kasir-pintar-ionpaytest/core"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Inquiry struct {
	core.Controller
}

func (a *Inquiry) Show(c *gin.Context) {

	data, err := io.ReadAll(c.Request.Body)
	// fmt.Println(data)

	url := "https://dev.nicepay.co.id/nicepay/direct/v2/inquiry"
	method := "POST"
	payload := bytes.NewReader(data)
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
	// fmt.Println(string(body))
	var success Success
	json.Unmarshal([]byte(body), &success)

	// Logging
	runLogFile, _ := os.OpenFile(
		"./log/inquiry.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	log.Logger = zerolog.New(multi)
	log.Info().
		Timestamp().
		Interface(success.ResultMsg, Success(success)).Msg("")

	c.JSON(http.StatusOK, success)
}
