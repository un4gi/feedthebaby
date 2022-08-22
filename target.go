package main

import (
	"encoding/json"
	"fmt"
)

const (
	TARGET_AUTH_CODE_URL   = "https://gsp.target.com/gsp/authentications/v1/auth_codes?client_id=ecom-web-1.0.0&redirect_uri=https%3A%2F%2Fwww.target.com%2F&acr=create_session_signin&state="
	TARGET_FULFILLMENT_URL = "https://redsky.target.com/redsky_aggregations/v1/web_platform/product_fulfillment_v1?key=9f36aeafbe60771e321a7cc95a78140772ab3e96&tcin="
	TARGET_HOMEPAGE_URL    = "https://www.target.com/"
	TARGET_LOGIN_URL       = ""
	TARGET_PRODUCT_URL     = "https://www.target.com/p/similac-pro-advance-non-gmo-powder-infant-formula/-/A-84185025?preselect="
)

type Product struct {
	Name    string
	URL     string
	Referer string
}

type TargetAvailability struct {
	Data struct {
		Product struct {
			TypeName    string `json:"__typename"`
			TCIN        string `json:"tcin"`
			Fulfillment struct {
				ProductID            string `json:"product_id"`
				OutOfStockEverywhere string `json:"is_out_of_stock_in_all_store_locations"`
				ShippingOptions      struct {
					AvailabilityStatus         string   `json:"availability_status"`
					LoyaltyAvailabilityStatus  string   `json:"loyalty_availability_status"`
					AvailableToPromiseQuantity float32  `json:"available_to_promise_quantity"`
					ReasonCode                 string   `json:"reason_code"`
					Services                   []string `json:"services"`
				} `json:"shipping_options"`
			} `json:"fulfillment"`
		} `json:"product"`
	} `json:"data"`
}

func CheckTargetItems() {
	var avail TargetAvailability
	var p1 = Product{"Similac Pro-Advance Non-GMO Powder Infant Formula (30.8 oz)", TARGET_FULFILLMENT_URL + "70000038", TARGET_PRODUCT_URL + "70000038"}
	var p2 = Product{"Similac Pro-Advance Non-GMO Powder Infant Formula (123.2 oz)", TARGET_FULFILLMENT_URL + "81624906", TARGET_PRODUCT_URL + "81624906"}

	bodyBytes := DoGetRequest(p1.URL, p1.Referer)
	_ = json.Unmarshal(bodyBytes, &avail)

	if avail.Data.Product.Fulfillment.ShippingOptions.AvailabilityStatus != "OUT_OF_STOCK" {
		SendDiscordMsg(&SendMsgData{
			Title:    fmt.Sprintf(p1.Name + " is available! Buy it before it is sold out! " + "(" + avail.Data.Product.Fulfillment.ShippingOptions.AvailabilityStatus + ")"),
			Url:      p1.Referer,
			Quantity: fmt.Sprintf("%g", avail.Data.Product.Fulfillment.ShippingOptions.AvailableToPromiseQuantity),
		})
	}

	bodyBytes = DoGetRequest(p2.URL, p2.Referer)
	_ = json.Unmarshal(bodyBytes, &avail)
	if avail.Data.Product.Fulfillment.ShippingOptions.AvailabilityStatus != "OUT_OF_STOCK" {
		SendDiscordMsg(&SendMsgData{
			Title:    fmt.Sprintf(p2.Name + " is available! Buy it before it is sold out! " + "(" + avail.Data.Product.Fulfillment.ShippingOptions.AvailabilityStatus + ")"),
			Url:      p2.Referer,
			Quantity: fmt.Sprintf("%g", avail.Data.Product.Fulfillment.ShippingOptions.AvailableToPromiseQuantity),
		})
	}
}
