package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/divyatambat/FarmersBasket/internal/app/user/mocks"
	"github.com/divyatambat/FarmersBasket/internal/pkg/apperrors"
	"github.com/divyatambat/FarmersBasket/internal/pkg/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProductHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Set up mock product service
		productSvc := new(mocks.Service)
		expectedProduct := dto.Product{ID: 1, Name: "Apple"}
		productSvc.On("GetProductByID", mock.Anything, int64(1)).Return(expectedProduct, nil)

		// Create handler instance
		handler := getProductHandler(productSvc)

		// Create request with valid product ID
		req, err := http.NewRequest(http.MethodGet, "/products/1", nil)
		assert.NoError(t, err)

		// Create response recorder
		recorder := httptest.NewRecorder()

		// Call the handler
		handler.ServeHTTP(recorder, req)

		// Assert response status and body
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, expectedProduct, recorder.Body.Bytes())

		// Assert mock interactions
		productSvc.AssertExpectations(t)
	})

	t.Run("invalid product ID", func(t *testing.T) {
		// Set up mock product service
		productSvc := new(mocks.Service)

		// Create handler instance
		handler := getProductHandler(productSvc)

		// Create request with invalid product ID
		req, err := http.NewRequest(http.MethodGet, "/products/invalid", nil)
		assert.NoError(t, err)

		// Create response recorder
		recorder := httptest.NewRecorder()

		// Call the handler
		handler.ServeHTTP(recorder, req)

		// Assert response status and body
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("error fetching product", func(t *testing.T) {
		// Set up mock product service
		productSvc := new(mocks.Service)
		productSvc.On("GetProductByID", mock.Anything, int64(1)).Return(dto.Product{}, apperrors.ProductNotFound.Error)

		// handler instance
		handler := getProductHandler(productSvc)

		// request with valid product ID
		req, err := http.NewRequest(http.MethodGet, "/products/1", nil)
		assert.NoError(t, err)

		// response recorder
		recorder := httptest.NewRecorder()

		// Calling handler
		handler.ServeHTTP(recorder, req)

		// Assert response status and body
		assert.Equal(t, http.StatusNotFound, recorder.Code)
		assert.Contains(t, recorder.Body.String(), apperrors.ProductNotFound.Error)

		// Assert mock interactions
		productSvc.AssertExpectations(t)
	})
}
