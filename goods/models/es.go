package models

type ESGoods struct {
	ID          int32   `json:"id"`
	CategoryID  int32   `json:"category_id"`
	BrandID     int32   `json:"brand_id"`
	Name        string  `json:"name"`
	Brief       string  `json:"brief"`
	MarketPrice float64 `json:"market_price"`
	ShopPrice   float64 `json:"shop_price"`

	SoldNum  int32 `json:"sold_num"`
	FavNum   int32 `json:"fav_num"`
	ClickNum int32 `json:"click_num"`

	IsOnSale       bool `json:"is_on_sale"`
	IsFreeShipping bool `json:"is_free_shipping"`
	IsNew          bool `json:"is_new"`
	IsHot          bool `json:"is_hot"`
}

func (ESGoods) GetIndexName() string {
	return "goods"
}

func (ESGoods) GetMapping() string {
	return `
	{
		"mappings" : {
			"properties" : {
				"id" : {
					"type" : "integer"
				},
				"brand_id" : {
					"type" : "integer"
				},
				"category_id" : {
					"type" : "integer"
				},
				"name" : {
					"type" : "text",
					"analyzer":"ik_max_word"
				},
				"brief" : {
					"type" : "text",
					"analyzer":"ik_max_word"
				},
				"market_price" : {
					"type" : "float"
				},
				"shop_price" : {
					"type" : "float"
				},
				"sold_num" : {
					"type" : "long"
				},
				"click_num" : {
					"type" : "integer"
				},
				"fav_num" : {
					"type" : "integer"
				},
				"is_free_shipping" : {
					"type" : "boolean"
				},
				"is_hot" : {
					"type" : "boolean"
				},
				"is_new" : {
					"type" : "boolean"
				},
				"is_on_sale" : {
					"type" : "boolean"
				}
			}
		}
	}`
}
