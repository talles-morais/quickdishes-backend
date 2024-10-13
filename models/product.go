package models

import "gopkg.in/validator.v2"

type Product struct {
	ProductID   string  `json:"product_id" validate:"nonzero" gorm:"primaryKey"`
	Name        string  `json:"name"        validate:"nonzero"`
	Value       float64 `json:"value"       validate:"nonzero"`
	Description string  `json:"description" gorm:"type:text"`
}

func ValidateProduct(product *Product) error {
	if err := validator.Validate(product); err != nil {
		return err
	}
	return nil
}
