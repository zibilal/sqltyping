package sqltyping

import (
	"encoding/json"
	"testing"
	"time"
)

var jString = `{
    "event_id" : "8b4c1afe-09a4-4b5d-96c4-f858ebb0ca80",
    "reference" : "201904010228340063942910",
    "order_split_id" : "8310240099850931653",
    "event_type" : "Order Success",
    "aggregate_id" : "f8942353-d9f2-4ccc-904d-e6b8858d9983",
    "version" : 5,
    "payload" : {
        "customer_id" : 45499118,
        "campaign_counter_detail_id" : 0,
        "discount" : 0.0,
        "id" : "f8942353-d9f2-4ccc-904d-e6b8858d9983",
        "reference" : "201904010228340063942910",
        "order_split_id" : "8310240099850931653",
        "shipment_type" : 1,
        "shipping_cost" : 0.0,
        "status" : "Success",
        "subtotal_price" : 20530.0,
        "vendor_id" : 84,
        "order_items" : [
            {
                "attributes" : "{\"additional_customer_data\":[{\"title\":\"NOMOR METER\",\"value\":\"01117153534\"},{\"title\":\"NAMA\",\"value\":\"MU'AS\"},{\"title\":\"TARIF/DAYA\",\"value\":\"R1/450 VA\"}],\"admin_fee\":2000,\"customer_number\":\"441230622134\",\"data_return\":{\"country_calling_code\":62,\"customer_id\":\"01117153534\",\"message\":\"\",\"phone_number\":\"87865675513\",\"price\":20530,\"product_code\":\"PLN20K\",\"transaction_date\":\"2019-04-01 14:29:13\",\"transaction_id\":\"1112617704038182912\",\"transaction_status\":\"processing\"},\"partner_price\":20530,\"product_code\":\"PLN20K\",\"token\":\"09949729684526979060\",\"total_price\":22000,\"transaction_amount\":20000,\"transaction_reference_number\":\"1323010201\"}",
                "description" : "Token Listrik Rp.20.000",
                "item_image" : "https://static.kudo.co.id/grab-integration/product-group1528100959.png",
                "commission" : 0.0,
                "name" : "Token Listrik Rp20000 - 01117153534",
                "product_code" : "PLN20K_01117153534",
                "price" : 20530.0,
                "quantity" : 1,
                "reseller_price" : 20000.0,
                "item_id" : 27325548,
                "retail_price" : 20530.0
            }
        ],
        "agent_id" : "3651944",
        "nis" : "201712091593",
        "device_id" : "5faf4fb89db76898",
        "total_commission" : 0.0,
        "shipping_trx_id" : "",
        "channel" : "MOBILE",
        "total_price" : 20530.0,
        "payment_type" : "ovo",
        "merchant_trx_id" : "",
        "cart_id" : "f2e46a972c177ebea77adcdaa05157c5086f6205a966083477b315e704723a82",
        "vendor_name" : "Kudo BillPayment",
        "customer" : {
            "name" : "Irwan Saputra",
            "hp" : "087865675513",
            "email" : "irwanwanky31@gmail.com",
            "number" : "01117153534"
        },
        "client_version" : "110"
    }
}`

type OrderEventTest struct {
	ID           string                  `bson:"event_id" json:"event_id"`
	Reference    string                  `bson:"reference" json:"reference"`
	OrderSplitID string                  `bson:"order_split_id" json:"order_split_id"`
	EventType    string                  `bson:"event_type" json:"event_type"`
	AggregateID  string                  `bson:"aggregate_id" json:"aggregate_id"`
	CreatedAt    time.Time               `bson:"created_at" json:"created_at"`
	Version      uint64                  `bson:"version" json:"version"`
	Payload      *OrderTest            `bson:"payload" json:"payload"`
}

type OrderTest struct {
	CreatedAt               time.Time    `json:"created_at" bson:"created_at,omitempty"`
	ExpiryDate              time.Time    `json:"expiry_date" bson:"expired_date,omitempty"`
	CustomerID              uint64       `json:"customer_id" bson:"customer_id"`
	CampaignCounterDetailID uint64       `json:"campaign_counter_detail_id" bson:"campaign_counter_detail_id"`
	Discount                float64      `json:"discount" bson:"discount"`
	ID                      string       `json:"id" bson:"id"`
	Reference               string       `json:"reference" bson:"reference"`
	OrderSplitID            string       `json:"order_split_id" bson:"order_split_id" transform:"order_split_id"`
	ShipmentType            int          `json:"shipment_type" bson:"shipment_type"`
	ShippingCost            float64      `json:"shipping_cost" bson:"shipping_cost"`
	Status                  string       `json:"status" bson:"status"`
	SubtotalPrice           float64      `json:"subtotal_price" bson:"subtotal_price" reflect:"total_paid"`
	VendorID                uint64       `json:"vendor_id" bson:"vendor_id"`
	OrderItems              []OrderItem2 `json:"order_items" bson:"order_items"`
	AgentID                 string       `json:"agent_id" bson:"agent_id" transform:"agent_id"`
	Nis                     string       `json:"nis" bson:"nis" transform:"nis"`
	DeviceID                string       `json:"device_id" bson:"device_id"`
	TotalCommission         float64      `json:"total_commission" bson:"total_commission"`
	ShippingTrxID           string       `json:"shipping_trx_id" bson:"shipping_trx_id"`
	Channel                 string       `json:"channel" bson:"channel"`
	TotalPrice              float64      `json:"total_price" bson:"total_price"`
	PaymentType             string       `json:"payment_type" bson:"payment_type" transform:"payment_type"`
	MerchantTrxID           string       `json:"merchant_trx_id" bson:"merchant_trx_id"`
	CartID                  string       `json:"cart_id" bson:"cart_id"`
	VendorName              string       `json:"vendor_name" bson:"vendor_name" transform:"vendor_name"`
	Customer                UserBasic    `json:"customer" bson:"customer" transform:"customer"`
	ClientVersion           string       `json:"client_version" bson:"client_version"`
}

type OrderItem2 struct {
	Attributes    string  `json:"attributes" bson:"attributes" transform:"attributes"`
	Description   string  `json:"description" bson:"description" transform:"description"`
	ItemImage     string  `json:"item_image" bson:"item_image" transform:"item_image"`
	Commission    float64 `json:"commission" bson:"commission" transform:"commission"`
	Name          string  `json:"name" bson:"name" transform:"item_name"`
	ProductCode   string  `json:"product_code" bson:"product_code" transform:"item_reference_id"`
	Price         float64 `json:"price" bson:"price" transform:"price"`
	Quantity      int     `json:"quantity" bson:"quantity" transform:"quantity"`
	ResellerPrice float64 `json:"reseller_price" bson:"reseller_price"`
	ItemID        uint64  `json:"item_id" bson:"item_id" transform:"item_id"`
	RetailPrice   float64 `json:"retail_price" bson:"retail_price" transform:"retail_price"`
}

type UserBasic struct {
	Name   string `json:"name" bson:"name" transform:"name"`
	Hp     string `json:"hp" bson:"hp" transform:"phone_number"`
	Email  string `json:"email" bson:"email" transform:"email"`
	Number string `json:"number" bson:"number" transform:"number"`
}

func TestQuery(t *testing.T) {
	t.Log("Test insert query 1")
	{
		mOrderEvent := OrderEventTest{}
		err := json.Unmarshal([]byte(jString), &mOrderEvent)
		if err != nil {
			t.Fatalf("%s expected err nil, got %s", failed, err.Error())
		}

		sTyping := NewSqlTyping(InsertQuery)
		results, err := sTyping.Typing(mOrderEvent)
		if err != nil {
			t.Fatalf("%s expected err nil, got %s", failed, err.Error())
		}

		t.Log("RESULT", results)
	}
}
