package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"kasir-pintar-ionpaytest/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payment struct {
	core.Controller
}

type Message struct {
	ResultCd       string
	ResultMsg      string
	TXid           string
	ReferenceNo    string
	PayMethod      string
	Amt            string
	TransDt        string
	TransTm        string
	Description    string
	AuthNo         string
	IssuBankCd     string
	AcquBankCd     string
	CardNo         string
	ReceiptCode    string
	MitraCd        string
	RecurringToken string
	PreauthToken   string
	Currency       string
	GoodsNm        string
	BillingNm      string
	CcTransType    string
	MRefNo         string
	InstmntType    string
	InstmntMon     string
	CardExpYymm    string
	IssuBankNm     string
	AcquBankNm     string
	TimeStamp      string
	MerchantToken  string
}

type Param struct {
	Amt           string `json:"amt"`
	MerchantToken string `json:"merchantToken"`
	CallBackUrl   string `json:"callBackUrl"`
	CardCvv       string `json:"cardCvv"`
	CardHolderNm  string `json:"cardHolderNm"`
	CardNo        string `json:"cardNo"`
	TimeStamp     string `json:"timeStamp"`
	CardExpYymm   string `json:"cardExpYymm"`
	TXid          string `json:"tXid"`
}

func (a *Payment) Show(c *gin.Context) {
	// params := c.Request.URL.Query()
	// var (params.Get("tXid"))
	var param Param
	param.Amt = c.Query("amt")
	param.MerchantToken = c.Query("merchantToken")
	param.CallBackUrl = c.Query("callBackUrl")
	param.CardCvv = c.Query("cardCvv")
	param.CardHolderNm = c.Query("cardHolderNm")
	param.CardNo = c.Query("cardNo")
	param.TimeStamp = c.Query("timeStamp")
	param.CardExpYymm = c.Query("cardExpYymm")
	param.TXid = c.Query("tXid")

	reqBytes, _ := json.Marshal(param)
	// fmt.Println("MASUK", string(reqBytes))
	fmt.Println("MASUK", string(reqBytes))

	// Write Request Log
	writeLog("request", string(reqBytes), "reqPayment")

	url := "https://dev.nicepay.co.id/nicepay/direct/v2/payment?timeStamp=" + param.TimeStamp + "&tXid=" + param.TXid + "&merchantToken=" + param.MerchantToken + "&amt=" + param.Amt + "&callBackUrl=" + param.CallBackUrl + "&cardNo=4222222222222222&cardExpYymm=" + param.CardExpYymm

	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "JSESSIONID=87039EA8196223F02AFF5663C012A171.pg_was1")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	// b, err := httputil.DumpResponse(res, true)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(string(b))
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))

	// var message Message
	// success
	// json.Unmarshal([]byte(body), &success)

	// // Logging
	// runLogFile, _ := os.OpenFile(
	// 	"./log/payment.log",
	// 	os.O_APPEND|os.O_CREATE|os.O_WRONLY,
	// 	0664,
	// )
	// multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	// log.Logger = zerolog.New(multi)
	// log.Info().
	// 	Timestamp().
	// 	Interface(success.ResultMsg, Success(success)).Msg("")
	// c.HTML(http.StatusOK, "response.html", res)

	// rspHtml := mahonia.NewDecoder("UTF-8").ConvertString(string(body))

	// doc, err := goquery.NewDocumentFromReader(strings.NewReader(rspHtml))
	// if err != nil {
	// 	log.Fatal().Msg("error")
	// }

	// inp := doc.Find("input")
	// v := inp.AttrOr("value", "")
	// message.AcquBankCd = v

	// getValue := func(name string) string {
	// 	// doc.Find("form").Each(func(i int, s *goquery.Selection) {
	// 	// 	s.Find("input").Each(func(i int, selection *goquery.Selection) {
	// 	// 		fmt.Printf("Review %d:  %s\n", i, selection.AttrOr("value", ""))
	// 	// 	})
	// 	// })
	// 	// fmt.Println(inp)
	// 	var hasil string
	// 	doc.Find("post").Each(func(i int, s *goquery.Selection) {
	// 		band, ok := s.Attr("href")
	// 		if ok {
	// 			fmt.Printf(band)
	// 			hasil = band
	// 			return
	// 		}
	// 	})
	// 	// value := doc.Find(name)
	// 	return hasil
	// }

	// // getValue("callbackUrl")
	// // fmt.Println(getValue("callbackUrl"))
	// jsonData := &Message{
	// 	ResultCd: getValue("ressd2ultCd"),
	// }

	// // header
	// var header string
	// for h, v := range res.Header {
	// 	for _, v := range v {
	// 		header += fmt.Sprintf("%s %s \n", h, v)
	// 	}
	// }

	// //append all to one slice
	// var write []byte
	// write = append(write, []byte(header)...)
	// write = append(write, body...)

	// write whole the body
	// write it to a file
	// err = ioutil.WriteFile("response.txt", body, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	// b, err := ioutil.ReadFile("response.txt")
	// if err != nil {
	// 	panic(err)
	// }
	writeLog("payment", string(body), "html")

	c.Data(http.StatusOK, "text/html; charset=utf-8", body)
	// c.JSON(http.StatusOK, html)
	// c.HTML(http.StatusOK, "markdown.html", gin.H{
	// 	"markdown": template.HTML(markdown),
	// })
}
