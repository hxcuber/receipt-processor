package model

import "time"

type Receipt struct {
	Retailer     string
	PurchaseDate time.Time
	PurchaseTime time.Time
	Items        []Item
	Total        float64
}
