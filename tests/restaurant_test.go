package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/talles-morais/quick-dishes/controllers"
	"github.com/talles-morais/quick-dishes/database"
	"github.com/talles-morais/quick-dishes/models"
	"github.com/talles-morais/quick-dishes/utils"
)

var restaurantID string

func CreateMockRestaurant() {
	hashedPassword, _ := utils.HashPassword("123456")

	restaurant := models.Restaurant{
		CNPJ:     "29276765000110",
		Name:     "Estabelecimento",
		Email:    "estab@lecimento.com",
		Password: hashedPassword,
		Address:  "Rua Principal, 123, Centro, Cidade - UF",
		Phone:    "22909090808",
	}
	database.DB.Create(&restaurant)
	restaurantID = restaurant.CNPJ
}

func DeleteMockRestaurant() {
	var restaurant models.Restaurant
	database.DB.Delete(&restaurant, restaurantID)
}

func TestCreateRestaurant(t *testing.T) {
	SetupDatabase()
	defer DeleteMockRestaurant()

	newRestaurant := models.Restaurant{
		CNPJ:     "29276765000110",
		Name:     "Estabelecimento",
		Email:    "estab@lecimento.com",
		Password: "123456",
		Address:  "Rua Principal, 123, Centro, Cidade - UF",
		Phone:    "22909090808",
	}
	restaurantID = newRestaurant.CNPJ

	jsonData, err := json.Marshal(newRestaurant)
	if err != nil {
		t.Fatalf("error converting to json: %v", err)
	}

	r := SetupRouter()
	r.POST("/signup", controllers.CreateRestaurant)

	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestLoginRestaurant(t *testing.T) {
	SetupDatabase()
	defer DeleteMockRestaurant()

	CreateMockRestaurant()

	loginData := struct {
		CNPJ     string `json:"cnpj"`
		Password string `json:"password"`
	}{
		CNPJ:     "29276765000110",
		Password: "123456",
	}

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		t.Fatalf("error converting to json: %v", err)
	}

	r := SetupRouter()
	r.POST("/login", controllers.LoginRestaurant)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
