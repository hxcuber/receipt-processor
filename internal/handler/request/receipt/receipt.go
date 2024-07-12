package receipt

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float32 `json:"price,string"`
}

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
	Total        float32 `json:"total,string"`
}
