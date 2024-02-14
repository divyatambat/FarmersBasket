package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/divyatambat/FarmersBasket/internal/app/product"
	"github.com/divyatambat/FarmersBasket/internal/pkg/apperrors"
	"github.com/divyatambat/FarmersBasket/internal/pkg/logger"
	"github.com/divyatambat/FarmersBasket/internal/pkg/middleware"
	"go.uber.org/zap"
)

func getProductHandler(productSvc product.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		rawProductID := vars["id"]

		productID, err := strconv.Atoi(rawProductID)
		if err != nil {
			logger.Errorw(ctx, "error occured while converting productID to an integer",
				zap.Error(err),
				zap.String("id", rawProductID),
			)

			middleware.ErrorResponse(ctx, w, http.StatusBadRequest, apperrors.ErrInvalidRequestParam)
			return
		}

		response, err := productSvc.GetProductByID(ctx, int64(productID))
		if err != nil {
			logger.Errorw(ctx, "error occured while fetching product info",
				zap.Error(err),
			)

			statusCode, errResponse := apperrors.MapError(err)
			middleware.ErrorResponse(ctx, w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(ctx, w, http.StatusOK, response)
	}
}

func listProductHandler(productSvc product.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		response, err := productSvc.ListProducts(ctx)
		if err != nil {
			logger.Errorw(ctx, "error occured while fetching product list",
				zap.Error(err),
			)

			middleware.ErrorResponse(ctx, w, http.StatusInternalServerError, apperrors.ErrInternalServerError)
			return
		}

		middleware.SuccessResponse(ctx, w, http.StatusOK, response)
	}
}

// func RegisterProductHandlers(router *mux.Router, productSvc product.Service) {
// 	// routes using Gorilla Mux syntax
// 	router.HandleFunc("/products/{id}", getProductHandler(productSvc))
// 	router.HandleFunc("/products", listProductHandler(productSvc))
// }
