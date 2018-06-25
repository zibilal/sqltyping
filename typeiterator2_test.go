package sqltyping

import (
	"encoding/json"
	"testing"
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
