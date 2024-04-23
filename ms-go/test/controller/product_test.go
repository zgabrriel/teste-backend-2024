package controller

import (
	"bytes"
	"encoding/json"
	"ms-go/app/models"
	"ms-go/router"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/likexian/gokit/assert"
)

func TestIndexProducts(t *testing.T) {

	body := gin.H{
		"produto": []models.Product{},
	}

	router := router.SetupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	response := struct {
		Message []models.Product `json:"data"`
	}{}

	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(body["produto"]), reflect.TypeOf(response.Message))

}

func TestShowProduct(t *testing.T) {
	body := gin.H{
		"produto": models.Product{},
	}

	router := router.SetupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/products/77", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	response := struct {
		Data models.Product `json:"data"`
	}{}

	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(body["produto"]), reflect.TypeOf(response.Data))
	assert.Equal(t, reflect.TypeOf(body["produto"]), reflect.TypeOf(response.Data))

}

func TestCreateProduct(t *testing.T) {
	body := gin.H{
		"produto": models.Product{},
	}

	router := router.SetupRouter()

	requestBody := []byte(`{
        "name": "teste",
        "brand": "teste",
        "price": 2999.99,
        "description": "Xbox + 1 Controles",
        "amount": 5
    }`)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	response := struct {
		Data models.Product `json:"data"`
	}{}

	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(body["produto"]), reflect.TypeOf(response.Data))
}

func TestUpdateProduct(t *testing.T) {
	body := gin.H{
		"produto": models.Product{},
	}

	router := router.SetupRouter()

	requestBody := []byte(`{
        "name": "teste do teste"
    }`)

	req, _ := http.NewRequest(http.MethodPatch, "/api/v1/products/133", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	response := struct {
		Data models.Product `json:"data"`
	}{}

	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(body["produto"]), reflect.TypeOf(response.Data))
}
