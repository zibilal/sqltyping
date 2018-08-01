package sqltyping

import (
	"encoding/json"
	"testing"
)

var strMsg = `{
	"add_partner_price": 0,
	"agent_distributor_id": 0,
	"attributes": "{\"customer\":{\"cellphone_number\":\"081220172375\"},\"purchase_referral\":{\"action\":\"pulsa\"}}",
	"cashier_id": null,
	"collection_id": 0,
	"created_at": "",
	"description": "tester ok",
	"device_id": null,
	"id": 0,
	"item_card_charging": 0,
	"item_delete_status": "",
	"item_discount": 0,
	"item_discount_shipping": 0,
	"item_image": "https://dev-static.kudo.co.id/api/images/category/ic_pulsa_telkomsel.png",
	"item_komisi": 0,
	"item_kudo_fee": 0,
	"item_name": "Voucher Telkomsel Rp5.000 (081220172375)",
	"item_reference_id": "78130379",
	"item_shipping": 0,
	"kudobox_id": null,
	"merchant_trx_id": null,
	"order_id": 252397162,
	"order_step": 0,
	"partner_plu": null,
	"price": 6000,
	"purchase_code": null,
	"qty_update_status": "",
	"quantity": 1,
	"refund_shipping_charges": 0,
	"reseller_price": 0,
	"retry_process": 0,
	"shipping_code": null,
	"shop_item_id": null,
	"sms": null,
	"spot_id": 0,
	"status": 0,
	"updated_at": "",
	"vendor_id": 83
}`

var strMsg2 = `{
	"add_partner_price": 0,
	"agent_distributor_id": 0,
	"attributes": "",
	"cashier_id": null,
	"collection_id": 0,
	"created_at": "",
	"description": "tester ok",
	"device_id": null,
	"id": 0,
	"item_card_charging": 0,
	"item_delete_status": "",
	"item_discount": 0,
	"item_discount_shipping": 0,
	"item_image": "https://dev-static.kudo.co.id/api/images/category/ic_pulsa_telkomsel.png",
	"item_komisi": 0,
	"item_kudo_fee": 0,
	"item_name": "Voucher Telkomsel Rp5.000 (081220172375)",
	"item_reference_id": "78130379",
	"item_shipping": 0,
	"kudobox_id": null,
	"merchant_trx_id": null,
	"order_id": 252397162,
	"order_step": 0,
	"partner_plu": null,
	"price": 6000,
	"purchase_code": null,
	"qty_update_status": "",
	"quantity": 1,
	"refund_shipping_charges": 0,
	"reseller_price": 0,
	"retry_process": 0,
	"shipping_code": null,
	"shop_item_id": null,
	"sms": null,
	"spot_id": 0,
	"status": 0,
	"updated_at": "",
	"vendor_id": 0
}`

func TestStructToStruct(t *testing.T) {
	t.Log("Testing struct to struct")
	{

		tmp := make(map[string]interface{})
		err := json.Unmarshal([]byte(strMsg), &tmp)

		if err != nil {
			t.Fatalf("%s 1. expected error nil, got %s", failed, err.Error())
		}

		if IsEmpty(tmp) {
			t.Fatalf("%s expected tmp not empty", failed)
		}

		coreItem := CoreOrderItems{}
		err = TypeIterator(tmp, &coreItem)
		if err != nil {
			t.Fatalf("%s 2. expected error nil, got %s", failed, err.Error())
		}

		if IsEmpty(coreItem) {
			t.Fatalf("%s expected coreItem not empty", failed)
		} else {
			b, _ := json.MarshalIndent(coreItem, "", "\t")
			t.Logf("%s expected coreItem not empty, got %s", success, string(b))
		}

		sqlTyping := NewSqlTyping(InsertQuery)
		queries, err := sqlTyping.Typing(coreItem)
		if err != nil {
			t.Fatalf("%s 3. expected error nil, got %s", failed, err.Error())
		}
		if len(queries) <= 0 {
			t.Fatalf("%s expected queries not empty, got %s", failed, "query is empty")
		} else {
			t.Logf("%s expected queries not empty, queris: %v", success, queries)
		}
	}

}
