package models

type Item struct {
	ItemID      uint   `json:"LineItemId" gorm:"primaryKey"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}
