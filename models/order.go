package models

import (
	"time"

	"gopkg.in/validator.v2"
)

type Order struct {
	ID          string     `json:"id"           validate:"nonzero"`
	Restaurant  Restaurant `json:"restaurant"   validate:"nonzero"`
	Status      string     `json:"status"       validate:"nonzero"`
	Pickup      bool       `json:"pickup"       validate:"nonzero"`
	OrderedAt   time.Time  `json:"ordered_at"   validate:"nonzero"`
	PreparedAt  time.Time  `json:"prepared_at"`
	DeliveredAt time.Time  `json:"delivered_at"`
	CompletedAt time.Time  `json:"completed_at"`
}

func ValidateOrder(order *Order) error {
	if err := validator.Validate(order); err != nil {
		return err
	}
	return nil
}
