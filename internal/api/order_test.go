package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/divyatambat/FarmersBasket/internal/app/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type orderInfo struct {
	ID           uint
	CustomerID   uint
	Order_Status string
}

func TestCreateOrderHandler(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		orderSvc := new(mocks.Service)
		orderInfo := &orderInfo{
			ID:           1,
			CustomerID:   1,
			Order_Status: "Placed",
		}
		orderSvc.On("CreateOrder", mock.Anything, mock.AnythingOfType("*order.OrderInfo")).Return(orderInfo, nil)
		req := httptest.NewRequest(http.MethodPost, "/orders", nil)

		w := httptest.NewRecorder()
		createOrderHandler(orderSvc)(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		orderSvc.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		orderSvc := new(mocks.Service)

		orderSvc.On("CreateOrder", mock.Anything, mock.AnythingOfType("*order.OrderInfo")).Return(nil, errors.New("error creating order"))
		req := httptest.NewRequest(http.MethodPost, "/orders", nil)

		w := httptest.NewRecorder()

		createOrderHandler(orderSvc)(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		orderSvc.AssertExpectations(t)

	})

}
