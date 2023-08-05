package services

import (
	"kasir-pintar-ionpaytest/core"
	"kasir-pintar-ionpaytest/models"

	"github.com/jinzhu/gorm"
)

type Registration struct {
	core.Service
}

type Response struct {
	Data string `json:"data"`
}

type Request struct {
	Data string `json:"data"`
}

func (s *Registration) All(registration interface{}) (*gorm.DB, []*models.Registration) {
	var list []*models.Registration
	result := s.Db().Where(registration).Find(&list)
	return result, list
}

func (s *Registration) Create(registration *models.Registration) (*gorm.DB, *models.Registration) {
	// DB
	result := s.Db().Create(&registration)

	// data, _ := json.Marshal(result.Value)
	// // fmt.Println(string(data))
	// url := "https://dev.nicepay.co.id/nicepay/direct/v2/registration"
	// method := "POST"
	// payload := bytes.NewReader(data)
	// payload := strings.NewReader(`{
	// 		"timeStamp":"20230805230817",
	// 		"iMid":"IONPAYTEST",
	// 		"payMethod":"01",
	// 		"currency":"IDR",
	// 		"amt":"100",
	// 		"referenceNo":"ord20230805230817",
	// 		"goodsNm":"Test Transaction Nicepay",
	// 		"billingNm":"Kasir Pintar User Test",
	// 		"billingPhone":"085648617037",
	// 		"billingEmail":"career@kasirpintar.co.id",
	// 		"billingAddr":"Jl. Dr. Ir. H. Soekarno No.19 Kel, Medokan Semampir, Kec. Sukolilo",
	// 		"billingCity":"Surabaya",
	// 		"billingState":"Jawa Timur",
	// 		"billingPostCd":"60119",
	// 		"billingCountry":"Indonesia",
	// 		"deliveryNm":"dobleh@merchant.com",
	// 		"deliveryPhone":"12345678",
	// 		"deliveryAddr":"Jl. Dr. Ir. H. Soekarno No.19 Kel, Medokan Semampir, Kec. Sukolilo",
	// 		"deliveryCity":"Surabaya",
	// 		"deliveryState":"Jawa Timur",
	// 		"deliveryPostCd":"60119",
	// 		"deliveryCountry":"Indonesia",
	// 		"dbProcessUrl":"https://ptsv2.com/t/test-nicepay-v2",
	// 		"vat":"",
	// 		"fee":"",
	// 		"notaxAmt":"",
	// 		"description":"",
	// 		"merchantToken":"d3bf58df89740ce19dc3bb7513145b93d2de474c86f112fb6d99fd71e1ea6f70",
	// 		"reqDt":"",
	// 		"reqTm":"",
	// 		"reqDomain":"https://kasirpintar.co.id/",
	// 		"reqServerIP":"127.0.0.1",
	// 		"reqClientVer":"",
	// 		"userIP":"127.0.0.1",
	// 		"userSessionID":"697D6922C961070967D3BA1BA5699C2C",
	// 		"userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML,like Gecko) Chrome/60.0.3112.101 Safari/537.36",
	// 		"userLanguage":"ko-KR,en-US;q=0.8,ko;q=0.6,en;q=0.4",
	// 		"cartData":"{\"count\":1,\"item\":[{\"img_url\":\"https:\/\/images.ctfassets.net\/od02wyo8cgm5\/14Njym0dRLAHaVJywF8WFL\/1910357dd0da0ae38b61bc8ad3db86e4\/cloudflyer_2-fw19-grey_lime-m-g1.png\",\"goods_name\":\"Shoe\",\"goods_detail\":\"Shoe\",\"goods_amt\":{{amt}}}]}",
	// 		"instmntType":"2",
	// 		"instmntMon":"1",
	// 		"recurrOpt":"0",
	// 		"bankCd":"CENA",
	// 		"vacctValidDt":"",
	// 		"vacctValidTm":"",
	// 		"merFixAcctId":"",
	// 		"mitraCd":""
	//   }`)

	// client := &http.Client{}
	// req, err := http.NewRequest(method, url, payload)

	// if err != nil {
	// 	fmt.Println(err)
	// 	// return
	// }
	// req.Header.Add("Content-Type", "application/json")

	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	// return
	// }
	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	// return
	// }
	// fmt.Println(string(body))
	return result, registration
}
