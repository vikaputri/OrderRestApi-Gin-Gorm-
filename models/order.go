package models

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderID      uint      `json:"orderId" gorm:"primaryKey "`
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
	OrderAt      time.Time `json:"orderAt"`
}

type FormOrder struct {
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []Item    `json:"items"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.OrderAt.IsZero() {
		o.OrderAt = time.Now()
	}

	errorMsg := ""

	if len(o.CustomerName) == 0 {
		errorMsg = errorMsg + "customerName can't be empty. "
	}

	for i, _ := range o.Items {
		if len(o.Items[i].ItemCode) == 0 {
			errorMsg = errorMsg + "itemCode item #" + strconv.Itoa(i+1) + " can't be empty. "
		}
	}

	if errorMsg != "" {
		err = errors.New(errorMsg)
	}

	return
}
