package models

import "gopkg.in/validator.v2"

type Client struct {
	ClientID string `json:"client_id" validate:"nonzero" gorm:"primaryKey"`
	Name     string `json:"name"      validate:"nonzero"`
	Address  string `json:"address"   validate:"nonzero"`
	Phone    string `json:"phone"     validate:"nonzero, regexp=^[0-9]*$"`
}

func ValidateClient(client *Client) error {
	if err := validator.Validate(client); err != nil {
		return err
	}
	return nil
}
