package sqltyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

type APIResponse struct {
	HTTPCode int         `json:"-"`
	Code     int         `json:"code"`
	Message  interface{} `json:"message"`
	Data     interface{} `json:"data,omitempty"`
}

type makePaymentResponse struct {
	Customer        customerMakePaymentResponse `json:"customer"`
	Items           []itemMakePaymentResponse   `json:"order_items"`
	OrderDate       string                      `json:"order_date"`
	PaymentInfo     interface{}                 `json:"payment_info"`
	Reference       string                      `json:"reference"`
	ShippingCost    string                      `json:"shipping_cost"`
	TotalCommission int                         `json:"total_commission"`
	TotalPaid       int                         `json:"total_paid"`
}

type itemMakePaymentResponse struct {
	Commission string `json:"item_komisi"`
	ItemImage  string `json:"item_image"`
	Name       string `json:"item_name"`
	Price      string `json:"price"`
	Quantity   string `json:"quantity"`
}

type customerMakePaymentResponse struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Phonenumber string `json:"phonenumber"`
}

func TestTypeIteratorWitNestedInterfaceStruct(t *testing.T) {
	t.Log("Test TypeIterator with nested interface struct")
	{
		sRef := struct {
			Message struct {
				Reference string `json:"reference"`
			} `json:"message"`
		}{struct {
			Reference string `json:"reference"`
		}{
			"2018112307070200253...",
		},
		}

		resp := APIResponse{
			Message: makePaymentResponse{
				Reference: "2018112307070200253",
			},
		}

		TypeIterator(sRef, &resp)

		t.Log(resp)
	}
}

type Input1 struct {
	Data string `json:"data"`
}

type Input2 struct {
	Data int `json:"data"`
}

type Output1 struct {
	Data interface{} `json:"data"`
}

func TestTypeIteratorInterfaceOutput(t *testing.T) {
	t.Log("Test TypeIteratorInterfaceOutput")
	{
		input1 := Input1{
			"2",
		}


		output1 := Output1{}
		TypeIterator(input1, &output1)

		if !reflect.DeepEqual(output1.Data, input1.Data) {
			b, _ := json.MarshalIndent(output1, "", "\t")
			t.Errorf("%s Unexpected output: %s", failed, string(b))
		}

		input2 := Input2{
			16,
		}
		output2 := Output1{}
		TypeIterator(input2, &output2)
		if !reflect.DeepEqual(output2.Data, input2.Data) {
			b, _ := json.MarshalIndent(output2, "", "\t")
			t.Errorf("%s Unexpected output: %s", failed, string(b))
		}
	}
}


var str = `
{"code":1000,"message":{"orders":[{"order_id":"372972980","order_split_id":"2910744890407510947","reference":"201906201126170084405210","total_paid":25000,"status_trx":"12","vendor_id":"83","vendor_name":"Kudo Airtime","vendor_type_id":"6","shipping_details":null,"created_at":"2019-06-20 11:26:17","order_items":[{"item_name":"Voucher Rp25.000 - 081337896767","status":6,"vendor_name":"Kudo Airtime"}],"ecommerce_order_status":0,"ecommerce_order_message":""},{"order_id":"372972967","order_split_id":"1549247040056951676","reference":"201906190602590019404170","total_paid":110000,"status_trx":"6","vendor_id":"86","vendor_name":"Kirim Uang","vendor_type_id":"6","shipping_details":null,"created_at":"2019-06-19 18:02:59","order_items":[{"item_name":"Kirim Uang\nBank BNI\nBpk ANDI  ANDI 111122223333","status":"6","vendor_name":"Kirim Uang"}],"ecommerce_order_status":0,"ecommerce_order_message":""},{"order_id":"372972964","order_split_id":"7089222747672602942","reference":"201906190550100033078781","total_paid":56500,"status_trx":"6","vendor_id":"86","vendor_name":"Kirim Uang","vendor_type_id":"6","shipping_details":null,"created_at":"2019-06-19 17:50:10","order_items":[{"item_name":"Kirim Uang\nBANK CENTRAL ASIA\nEVI SURYANINGSIH 1111","status":"6","vendor_name":"Kirim Uang"}],"ecommerce_order_status":0,"ecommerce_order_message":""},{"order_id":"372972948","order_split_id":"15609295996327","reference":"20190619024","total_paid":135313,"status_trx":"3","vendor_id":"16","vendor_name":"Makna Karsa Mulya","vendor_type_id":"5","shipping_details":null,"created_at":"2019-06-19 14:33:19","order_items":[{"item_name":"BPJS Kesehatan 1 Bulan (No.Pelanggan: 8888801302479201)","status":"6","vendor_name":"Makna Karsa Mulya"}],"ecommerce_order_status":0,"ecommerce_order_message":""},{"order_id":"372972945","order_split_id":"4541790411466804856","reference":"201906190155340070646645","total_paid":52080,"status_trx":"3","vendor_id":"84","vendor_name":"Kudo Billpayment","vendor_type_id":"5","shipping_details":null,"created_at":"2019-06-19 13:55:34","order_items":[{"item_name":"Tagihan Listrik - 123123123123","status":"3","vendor_name":"Kudo Billpayment"}],"ecommerce_order_status":0,"ecommerce_order_message":""},{"order_id":"372972926","order_split_id":"9484338928916035108","reference":"201906191136340084790544","total_paid":25000,"status_trx":"3","vendor_id":"83","vendor_name":"Kudo Airtime","vendor_type_id":"6","shipping_details":null,"created_at":"2019-06-19 11:36:34","order_items":[{"item_name":"Voucher Rp25.000 - 081339372389","status":"3","vendor_name":"Kudo Airtime"}],"ecommerce_order_status":0,"ecommerce_order_message":""},{"order_id":"372972874","order_split_id":"9917427783336624839","reference":"201906180536250077427848","total_paid":56500,"status_trx":"6","vendor_id":"86","vendor_name":"Kirim Uang","vendor_type_id":"6","shipping_details":null,"created_at":"2019-06-18 17:36:26","order_items":[{"item_name":"Kirim Uang\nBANK CENTRAL ASIA\nEVI SURYANINGSIH 11112222333","status":"6","vendor_name":"Kirim Uang"}],"ecommerce_order_status":0,"ecommerce_order_message":""},{"order_id":"372972871","order_split_id":"3692621314427666003","reference":"201906180527260067672867","total_paid":10600,"status_trx":"3","vendor_id":"87","vendor_name":"Isi Saldo OVO","vendor_type_id":"6","shipping_details":null,"created_at":"2019-06-18 17:27:27","order_items":[{"item_name":"OVO Top Up Nominal Rp10.000 - 085735104171","status":"3","vendor_name":"Isi Saldo OVO"}],"ecommerce_order_status":0,"ecommerce_order_message":""},{"order_id":"372972865","order_split_id":"3816884510688477597","reference":"201906180518200051332760","total_paid":25000,"status_trx":"3","vendor_id":"83","vendor_name":"Kudo Airtime","vendor_type_id":"6","shipping_details":null,"created_at":"2019-06-18 17:18:21","order_items":[{"item_name":"Voucher Rp25.000 - 082234413346","status":"4","vendor_name":"Kudo Airtime"}],"ecommerce_order_status":0,"ecommerce_order_message":""},{"order_id":"372972694","order_split_id":"15607578085205","reference":"20190617023","total_paid":84167,"status_trx":"3","vendor_id":"72","vendor_name":"Toko Kudo","vendor_type_id":"3","shipping_details":{"courier":"REGULER + ASURANSI","insurance_cost":"5078","shipping_detail":"","shipping_cost_vendor":47078,"shipping_cost":47078,"save":true,"shipping_type":0,"weight":50,"address":"simpang agung","city":"Kab. Lampung Tengah","recipient_name":"ika","kecamatan":"Seputih Agung","province":"Lampung","email":"nurzanah@ymail.com","kelurahan":"Simpang Agung","phone_number":"081234567890","ownership":false,"kelurahan_id":72322,"kecamatan_id":1805081,"province_id":18,"city_id":1805,"address_id":721731,"postcode":34166,"status_shipping":null,"shipping_code":"None","shipping_history":[{"status":"Pesanan Anda sedang dalam proses. Status pengiriman akan kami informasikan kembali","date":"2019-06-17 14:50:08"}],"status":1,"message":"Dipesan"},"created_at":"2019-06-17 14:50:08","order_items":[{"item_name":"SoftCase Jelly Vivo- Anti Crack Vivo","status":"6","vendor_name":"Toko Kudo"}],"ecommerce_order_status":1,"ecommerce_order_message":"Dipesan"}]}}
`


type VersionThreeBaseResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

type DataOrders struct {
	Orders []OrderDetails `json:"orders"`
}

type OrderDetails struct {
	OrderId               string      `json:"order_id"`
	OrderSplitId          string      `json:"order_split_id"`
	Reference             string      `json:"reference"`
	TotalPaid             int64       `json:"total_paid"`
	StatusTrx             string      `json:"status_trx"`
	VendorId              string      `json:"vendor_id"`
	VendorName            string      `json:"vendor_name"`
	VendorTypeId          string      `json:"vendor_type_id"`
	ShippingDetails       interface{} `json:"shipping_details"`
	CreatedAt             string      `json:"created_at"`
	OrderItems            []Items     `json:"order_items"`
	EcommerceOrderStatus  int         `json:"ecommerce_order_status"`
	EcommerceOrderMessage string      `json:"ecommerce_order_message"`
}

type Items struct {
	ItemName   string      `json:"item_name"`
	Status     interface{} `json:"status"`
	VendorName string      `json:"vendor_name"`
}

func TestTypeIteratorOrderDetails(t *testing.T) {
	t.Log("Test TypeIteratorOrderDetails")
	{
		response := VersionThreeBaseResponse{}
		err := json.Unmarshal([]byte(str), &response)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		}

		dataOrders := DataOrders{}

		TypeIterator(response.Message, &dataOrders)

		b, _ := json.MarshalIndent(dataOrders, "", "\t")
		t.Log(string(b))
	}
}