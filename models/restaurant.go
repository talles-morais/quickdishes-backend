package models

type Restaurant struct {
	CNPJ     string `json:"cnpj"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}
