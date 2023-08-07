package models

type Inquiry struct {
	TimeStamp     string `json:"timeStamp"`
	TXid          string `json:"tXid"`
	IMid          string `json:"iMid"`
	ReferenceNo   string `json:"referenceNo"`
	Amt           string `json:"amt"`
	MerchantToken string `json:"merchantToken"`
}
