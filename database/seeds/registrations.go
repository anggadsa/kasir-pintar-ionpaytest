package seeds

import (
	"kasir-pintar-ionpaytest/config"
	"kasir-pintar-ionpaytest/models"

	"log"
)

func Registrations() {
	db := config.Db()
	// var Registrations []models.Registration

	// result := db.Find(&Registrations)
	// if result.RowsAffected >= 4 {
	// 	log.Printf("Seed (Registritions): Nothing to seed")
	// 	return
	// }

	Registrations := models.Registration{
		TimeStamp:       "2018012310j0505",
		IMid:            "IONPAYTEST",
		PayMethod:       "01",
		Currency:        "IDR",
		Amt:             "10000",
		ReferenceNo:     "ORD12345",
		GoodsNm:         "Test Transaction Nicepay",
		BillingNm:       "John Doe",
		BillingPhone:    "12345678",
		BillingEmail:    "email@merchant.com",
		BillingAddr:     "Jalan Bukit Berbunga 22",
		BillingCity:     "Jakarta",
		BillingState:    "DKI Jakarta",
		BillingPostCd:   "12345",
		BillingCountry:  "Indonesia",
		DeliveryNm:      "email@merchant.com",
		DeliveryPhone:   "12345678",
		DeliveryAddr:    "Jalan Bukit Berbunga 22",
		DeliveryCity:    "Jakarta",
		DeliveryState:   "DKI Jakarta",
		DeliveryPostCd:  "12345",
		DeliveryCountry: "Indonesia",
		DbProcessUrl:    "https://merchant.com/api/dbProcessUrl/Notif",
		MerchantToken:   "f9d30f6c972e2b5718751bd087b178534673a91bbac845f8a24e60e8e4abbbc5",
		ReqDomain:       "merchant.com",
		ReqServerIP:     "127.0.0.1",
		UserIP:          "127.0.0.1",
		UserSessionID:   "697D6922C961070967D3BA1BA5699C2C",
		UserAgent:       "Mozilla/5.0 (Windows NT 10.0; Win64; X64) AppleWebKit/537.36 (KHTML,like Gecko) Chrome/60.0.3112.101 Safari/537.36",
		UserLanguage:    "ko-KR,en-US;q=0.8,ko;q=0.6,en;q=0.4",
		InstmntType:     "2",
		InstmntMon:      "1",
		RecurrOpt:       "2",
	}
	db.Create(&Registrations)
	log.Printf("Seed (Perfils): Success")
}
