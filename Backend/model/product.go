package model

import "time"

type Product struct {
	ID           int        `json:"id"`
	Product_name string     `json:"product_name"`
	SKU          string     `json:"sku"`
	Quantity     int        `json:"quantity"`
	Status       string     `json:"status"`
	Location     string     `json:"location"`
	Created_at   *time.Time `json:"created_at"`
	Updated_at   *time.Time `json:"updated_at"`
}
