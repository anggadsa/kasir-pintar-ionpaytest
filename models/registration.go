package models

type Registration struct {
	TimeStamp       string `json:"timeStamp" gorm:"type:numeric"`
	IMid            string `json:"iMid"`
	PayMethod       string `json:"payMethod"`
	Currency        string `json:"currency"`
	Amt             string `json:"amt"`
	ReferenceNo     string `json:"referenceNo"`
	GoodsNm         string `json:"goodsNm"`
	BillingNm       string `json:"billingNm"`
	BillingPhone    string `json:"billingPhone"`
	BillingEmail    string `json:"billingEmail"`
	BillingAddr     string `json:"billingAddr"`
	BillingCity     string `json:"billingCity"`
	BillingState    string `json:"billingState"`
	BillingPostCd   string `json:"billingPostCd"`
	BillingCountry  string `json:"billingCountry"`
	DeliveryNm      string `json:"deliveryNm"`
	DeliveryPhone   string `json:"deliveryPhone"`
	DeliveryAddr    string `json:"deliveryAddr"`
	DeliveryCity    string `json:"deliveryCity"`
	DeliveryPostCd  string `json:"deliveryPostCd"`
	DeliveryState   string `json:"deliveryState"`
	DeliveryCountry string `json:"deliveryCountry"`
	DbProcessUrl    string `json:"dbProcessUrl"`
	MerchantToken   string `json:"merchantToken"`
	ReqDomain       string `json:"reqDomain"`
	ReqServerIP     string `json:"reqServerIP"`
	UserIP          string `json:"userIP"`
	UserSessionID   string `json:"userSessionID"`
	UserAgent       string `json:"userAgent"`
	UserLanguage    string `json:"userLanguage"`
	InstmntType     string `json:"instmntType"`
	InstmntMon      string `json:"instmntMon"`
	RecurrOpt       string `json:"recurrOpt"`
}

// Read implements io.Reader.
func (*Registration) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}
