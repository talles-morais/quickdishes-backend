package models

import (
	"time"

	"gopkg.in/validator.v2"
)

type Order struct {
	OrderID     string    `json:"order_id"     validate:"nonzero" gorm:"primaryKey"`
	Restaurant  string    `json:"restaurant"   validate:"nonzero" gorm:"foreignKey:CNPJ"`
	Products    []Product `json:"products"     gorm:"many2many:order_products;joinForeignKey:OrderID;joinReferences:ProductID;constraint:OnDelete:CASCADE;"`
	Status      string    `json:"status"       validate:"nonzero"`
	Pickup      bool      `json:"pickup"`
	OrderedAt   time.Time `json:"ordered_at"   validate:"nonzero"`
	PreparedAt  time.Time `json:"prepared_at"`
	DeliveredAt time.Time `json:"delivered_at"`
	CompletedAt time.Time `json:"completed_at"`
}

func ValidateOrder(order *Order) error {
	if err := validator.Validate(order); err != nil {
		return err
	}
	return nil
}
