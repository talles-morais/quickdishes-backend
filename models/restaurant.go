package models

import (
	"gopkg.in/validator.v2"
)

type Restaurant struct {
	CNPJ     string `json:"cnpj"     validate:"nonzero, len=14"  gorm:"primaryKey"`
	Name     string `json:"name"     validate:"nonzero"`
	Email    string `json:"email"    validate:"nonzero, regexp=^([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,})$"`
	Password string `json:"password" validate:"nonzero"`
	Address  string `json:"address"  validate:"nonzero"`
	Phone    string `json:"phone"    validate:"nonzero, regexp=^[0-9]*$"`
}

func ValidateRestaurant(restaurant *Restaurant) error {
	if err := validator.Validate(restaurant); err != nil {
		return err
	}
	return nil
}
