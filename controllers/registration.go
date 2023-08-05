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

	"github.com/gin-gonic/gin"
)

type Registration struct {
	core.Controller
}

type Success struct {
	ResultCd     string
	ResultMsg    string
	TXid         string
	ReferenceNo  string
	PayMethod    string
	Amt          string
	TransDt      string
	TransTm      string
	Description  string
	BankCd       string
	VacctNo      string
	MitraCd      string
	PayNo        string
	Currency     string
	GoodsNm      string
	BillingNm    string
	VacctValidDt string
	VacctValidTm string
	PayValidDt   string
	PayValidTm   string
	RequestURL   string
	PaymentExpDt string
	PaymentExpTm string
	QrContent    string
	QrUrl        string
}

func (a *Registration) Index(c *gin.Context) {
	var params *models.Registration
	c.BindJSON(&params)
	// c.ShouldBind(&params)
	result, list := new(services.Registration).All(params)
	a.Json(c, result, list)
}

func (a *Registration) Store(c *gin.Context) {
	var params *models.Registration
	c.BindJSON(&params)

	result, _ := new(services.Registration).Create(params)

	data, _ := json.Marshal(result.Value)
	// fmt.Println(string(data))
	url := "https://dev.nicepay.co.id/nicepay/direct/v2/registration"
	method := "POST"
	payload := bytes.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(body))

	var success Success
	json.Unmarshal([]byte(body), &success)
	a.Json(c, result, success)
}
