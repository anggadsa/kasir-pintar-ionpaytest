package controllers

import (
	"fmt"
	"html"
	"net/http"
)

type Index struct {
	// core.Controller
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

func (i *Index) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

}
