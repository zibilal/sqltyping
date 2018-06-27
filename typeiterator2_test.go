package sqltyping

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

type CoreOrderItems struct {
	AddPartnerPrice       int         `json:"add_partner_price"`
	AgentDistributorID    int         `json:"agent_distributor_id"`
	Attributes            string      `json:"attributes"`
	CashierID             interface{} `json:"cashier_id"`
	CollectionID          int         `json:"collection_id"`
	CreatedAt             string      `json:"created_at"`
	Description           string      `json:"description"`
	DeviceID              interface{} `json:"device_id"`
	ID                    int         `json:"id"`
	ItemCardCharging      int         `json:"item_card_charging"`
	ItemDeleteStatus      string      `json:"item_delete_status"`
	ItemDiscount          int         `json:"item_discount"`
	ItemDiscountShipping  int         `json:"item_discount_shipping"`
	ItemImage             string      `json:"item_image"`
	ItemKomisi            float64     `json:"item_komisi"`
	ItemKudoFee           int         `json:"item_kudo_fee"`
	ItemName              string      `json:"item_name" transform:"item_name"`
	ItemReferenceID       string      `json:"item_reference_id"`
	ItemShipping          int         `json:"item_shipping"`
	KudoboxID             interface{} `json:"kudobox_id"`
	MerchantTrxID         interface{} `json:"merchant_trx_id"`
	OrderID               uint64      `json:"order_id"`
	OrderStep             int         `json:"order_step"`
	PartnerPlu            interface{} `json:"partner_plu"`
	Price                 float32     `json:"price"`
	PurchaseCode          interface{} `json:"purchase_code"`
	QtyUpdateStatus       string      `json:"qty_update_status"`
	Quantity              int         `json:"quantity"`
	RefundShippingCharges int         `json:"refund_shipping_charges"`
	ResellerPrice         float64     `json:"reseller_price"`
	RetryProcess          int         `json:"retry_process"`
	ShippingCode          interface{} `json:"shipping_code"`
	ShopItemID            interface{} `json:"shop_item_id"`
	Sms                   interface{} `json:"sms"`
	SpotID                int         `json:"spot_id"`
	Status                int         `json:"status"`
	UpdatedAt             string      `json:"updated_at"`
	VendorID              int         `json:"vendor_id"`
}

type OrderItemTest struct {
	Attributes    string  `json:"attributes" bson:"attributes" transform:"attributes"`
	Description   string  `json:"description" bson:"description"`
	ItemImage     string  `json:"item_image" bson:"item_image" transform:"item_image"`
	Commission    float32 `json:"commission" bson:"commission" transform:"commission"`
	Name          string  `json:"name" bson:"name" transform:"item_name"`
	ProductCode   string  `json:"product_code" bson:"product_code" transform:"item_reference_id"`
	Price         float32 `json:"price" bson:"price" transform:"price"`
	Quantity      int     `json:"quantity" bson:"quantity" transform:"quantity"`
	ResellerPrice float64 `json:"reseller_price" bson:"reseller_price"`
	ItemID        uint64  `json:"item_id" bson:"item_id" transform:"item_id"`
}

type OrderEventTesting struct {
	ID          string        `bson:"event_id" json:"event_id"`
	Reference   string        `bson:"reference" json:"reference"`
	EventType   string        `bson:"event_type" json:"event_type"`
	AggregateID string        `bson:"aggregate_id" json:"aggregate_id"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updated_at"`
	Version     int           `bson:"version" json:"version"`
	Payload     *OrderTesting `bson:"payload" json:"payload"`
}

type OrderTesting struct {
	CreatedAt       time.Time     `json:"created_at" bson:"created_at"`
	ExpiryDate      time.Time     `json:"expiry_date" bson:"expired_date"`
	CustomerID      uint64        `json:"customer_id" bson:"customer_id"`
	CampaignID      string        `json:"campaign_id" bson:"campaign_id"`
	ID              string        `json:"id" bson:"id"`
	Reference       string        `json:"reference" bson:"reference"`
	ShipmentType    int           `json:"shipment_type" bson:"shipment_type"`
	ShippingCost    float32       `json:"shipping_cost" bson:"shipping_cost"`
	Status          string        `json:"status" bson:"status"`
	SubtotalPrice   float32       `json:"subtotal_price" bson:"subtotal_price"`
	UserID          uint64        `json:"user_id" bson:"user_id"`
	VendorID        uint64        `json:"vendor_id" bson:"vendor_id"`
	OrderItems      []OrderItemEx `json:"order_items" bson:"order_items"`
	AgentID         uint64        `json:"agent_id" bson:"agent_id"`
	DeviceID        string        `json:"device_id" bson:"device_id"`
	TotalCommission float32       `json:"total_commission" bson:"total_commission"`
	ShippingTrxID   string        `json:"shipping_trx_id" bson:"shipping_trx_id"`
	Channel         string        `json:"channel" bson:"channel"`
	TotalPrice      float32       `json:"total_price" bson:"total_price"`
	PaymentType     string        `json:"payment_type" bson:"payment_type"`
	MerchantTrxID   string        `json:"merchant_trx_id" bson:"merchant_trx_id"`
	CartID          string        `json:"cart_id" bson:"cart_id"`
}

type OrderItemEx struct {
	Attributes    string  `json:"attributes" bson:"attributes" transform:"attributes"`
	Description   string  `json:"description" bson:"description"`
	ItemImage     string  `json:"item_image" bson:"item_image" transform:"item_image"`
	Commission    float32 `json:"commission" bson:"commission" transform:"commission"`
	Name          string  `json:"name" bson:"name" transform:"item_name"`
	ProductCode   string  `json:"product_code" bson:"product_code" transform:"item_reference_id"`
	Price         float32 `json:"price" bson:"price" transform:"price"`
	Quantity      int     `json:"quantity" bson:"quantity" transform:"quantity"`
	ResellerPrice float64 `json:"reseller_price" bson:"reseller_price"`
	ItemID        uint64  `json:"item_id" bson:"item_id" transform:"item_id"`
}

func TestOrderItemIterate(t *testing.T) {
	t.Log("Test Order item")
	{
		orderItem := OrderItemTest{
			Attributes:    "The attributes",
			Description:   "The description",
			ItemImage:     "The Image",
			Commission:    12.31,
			Name:          "The Name",
			ProductCode:   "Product code",
			Quantity:      15,
			ResellerPrice: 15000,
			ItemID:        125,
		}

		coreItem := CoreOrderItems{}

		err := TypeIterator(orderItem, &coreItem)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
			if IsEmpty(coreItem) {
				t.Fatalf("%s expected core item not empty", failed)
				return
			}
			b, _ := json.MarshalIndent(coreItem, "", "\t")
			t.Logf("%s RESULT: %s", success, string(b))

		}
	}
}

type JsonConvertedMessage struct {
	Schema  interface{}
	Payload Payload
}

type Payload struct {
	After     interface{} `json:"after"`
	Before    interface{} `json:"before"`
	Patch     interface{} `json:"patch"`
	Op        interface{} `json:"op"`
	Timestamp int64       `json:"ts_ms"`
}

func TestOrderEventTesting(t *testing.T) {
	t.Log("Test OrderEvent testing")
	{
		testingTxt := `{"schema":{"type":"struct","fields":[{"type":"string","optional":true,"name":"io.debezium.data.Json","version":1,"field":"after"},{"type":"string","optional":true,"name":"io.debezium.data.Json","version":1,"field":"patch"},{"type":"struct","fields":[{"type":"string","optional":false,"field":"name"},{"type":"string","optional":false,"field":"rs"},{"type":"string","optional":false,"field":"ns"},{"type":"int32","optional":false,"field":"sec"},{"type":"int32","optional":false,"field":"ord"},{"type":"int64","optional":true,"field":"h"},{"type":"boolean","optional":true,"field":"initsync"}],"optional":false,"name":"io.debezium.connector.mongo.Source","version":1,"field":"source"},{"type":"string","optional":true,"field":"op"},{"type":"int64","optional":true,"field":"ts_ms"}],"optional":false,"name":"events_sourcing.terracotta.events.Envelope"},"payload":{"after":"{\"_id\" : {\"$oid\" : \"5b31100800cfd3e83d65bb59\"},\"event_id\" : \"f32ceec5-8e7d-43bd-ba66-1eabe30f4526\",\"reference\" : \"201806251053360082028569\",\"event_type\" : \"Initiated\",\"aggregate_id\" : \"cb2a0456-9859-4a9f-858b-b12cde004517\",\"created_at\" : {\"$date\" : 1529942016618},\"updated_at\" : {\"$date\" : 1529942016618},\"version\" : 1,\"payload\" : {\"created_at\" : {\"$date\" : 1529942016618},\"expired_date\" : {\"$date\" : -62135596800000},\"customer_id\" : {\"$numberLong\" : \"0\"},\"campaign_id\" : \"\",\"id\" : \"6f134af9-76cb-41ef-9036-3e00a8951d02\",\"reference\" : \"201806251053360082028569\",\"shipment_type\" : 0,\"shipping_cost\" : 0.0,\"status\" : \"Initiated\",\"subtotal_price\" : 0.0,\"user_id\" : {\"$numberLong\" : \"1234\"},\"vendor_id\" : {\"$numberLong\" : \"19\"},\"order_items\" : [{\"attributes\" : \"\",\"description\" : \"\",\"item_image\" : \"\",\"commission\" : 0.0,\"name\" : \"Pulsa 1 Test\",\"product_code\" : \"\",\"price\" : 0.0,\"quantity\" : 2,\"reseller_price\" : 0.0,\"item_id\" : {\"$numberLong\" : \"44321\"}}, {\"attributes\" : \"\",\"description\" : \"\",\"item_image\" : \"\",\"commission\" : 0.0,\"name\" : \"Pulsa 2 Test\",\"product_code\" : \"\",\"price\" : 0.0,\"quantity\" : 2,\"reseller_price\" : 0.0,\"item_id\" : {\"$numberLong\" : \"44321\"}}],\"agent_id\" : {\"$numberLong\" : \"0\"},\"device_id\" : \"\",\"total_commission\" : 0.0,\"shipping_trx_id\" : \"\",\"channel\" : \"\",\"total_price\" : 0.0,\"payment_type\" : \"Pulsa / PPOB Credit\",\"merchant_trx_id\" : \"\",\"cart_id\" : \"\"}}","patch":null,"source":{"name":"events-sourcing","rs":"rs-my-mongo-command","ns":"terracotta.events","sec":1529942024,"ord":2,"h":-5500174221298992661,"initsync":null},"op":"c","ts_ms":1529942024460}}`
		convertedMessage := JsonConvertedMessage{}
		err := json.Unmarshal([]byte(testingTxt), &convertedMessage)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
			if IsEmpty(convertedMessage) {
				t.Fatalf("%s expected convertedMessage is not empty", failed)
			} else {
				b, _ := json.MarshalIndent(convertedMessage, "", "\t")
				t.Logf("%s RESULT: %s", success, string(b))
			}
		}

		eventTesting := OrderEventTesting{}
		afterMap := make(map[string]interface{})
		json.Unmarshal([]byte(convertedMessage.Payload.After.(string)), &afterMap)

		handleDateTag := func(input interface{}) (interface{}, error) {
			if minput, ok := input.(map[string]interface{}); ok {
				d, found := minput["$date"]
				if found {
					if dfloat64, ok := d.(float64); ok && dfloat64 > 0 {
						dint64 := int64(dfloat64) / 1000

						return time.Unix(dint64, 0), nil
					}
				}
			}

			return nil, fmt.Errorf("unable to handle data: %v", input)
		}

		handleIdTag := func(input interface{}) (interface{}, error) {
			if minput, ok := input.(map[string]interface{}); ok {
				s, found := minput["$oid"]
				if found {
					if str, ok := s.(string); ok {
						return str, nil
					}
				}
			}

			return nil, fmt.Errorf("unable to handle data: %v", input)
		}

		handleLongTag := func(input interface{}) (interface{}, error) {
			if minput, ok := input.(map[string]interface{}); ok {
				d, found := minput["$numberLong"]
				if found {

					if dstring, ok := d.(string); ok {
						return strconv.ParseUint(dstring, 10, 64)
					}
				}
			}

			return nil, fmt.Errorf("unable to handle data: %v", input)
		}
		err = TypeIterator(afterMap, &eventTesting, handleDateTag, handleIdTag, handleLongTag)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)

			if IsEmpty(eventTesting) {
				t.Fatalf("%s expected convertedMessage is not empty", failed)
			} else {
				b, _ := json.MarshalIndent(eventTesting, "", "\t")
				t.Logf("%s RESULT: %s", success, string(b))
			}
		}
	}
}
