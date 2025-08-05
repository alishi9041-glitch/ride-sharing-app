package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"joi-delivery-golang/internal/dto/request"
	"joi-delivery-golang/internal/dto/response"
	"joi-delivery-golang/internal/models"
	"joi-delivery-golang/internal/service"
)

func setupTestServer() *Server {
	service.ClearTestData()
	ctx := context.Background()
	server := NewServer(ctx)
	return server
}

func TestAddToCart_Success_Store(t *testing.T) {
	server := setupTestServer()
	e := echo.New()

	reqBody := request.AddToCartRequest{
		UserID:    "user101",
		OutletID:  "store101",
		ProductID: "product101",
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/cart/product", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := server.handlers.AddToCart(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp response.AddToCartResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Cart)
}

func TestGetCart_Success(t *testing.T) {
	server := setupTestServer()
	e := echo.New()

	reqBody := request.AddToCartRequest{
		UserID:    "user101",
		OutletID:  "store101",
		ProductID: "product101",
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/cart/product", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := server.handlers.AddToCart(c)
	assert.NoError(t, err)

	var addResponse response.AddToCartResponse
	err = json.Unmarshal(rec.Body.Bytes(), &addResponse)
	assert.NoError(t, err)

	req = httptest.NewRequest(http.MethodGet, "/cart/view?userId="+addResponse.Cart.User.ID, nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	err = server.handlers.GetCart(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var cart models.Cart
	err = json.Unmarshal(rec.Body.Bytes(), &cart)
	assert.NoError(t, err)
	assert.Equal(t, "user101", cart.User.ID)
	assert.Len(t, cart.Products, 1)
	assert.Equal(t, "product101", cart.Products[0].ID)
}
