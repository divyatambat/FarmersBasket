package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/divyatambat/FarmersBasket/internal/app/order"
	"github.com/divyatambat/FarmersBasket/internal/pkg/apperrors"
	"github.com/divyatambat/FarmersBasket/internal/pkg/dto"
	"github.com/divyatambat/FarmersBasket/internal/pkg/logger"
	"github.com/divyatambat/FarmersBasket/internal/pkg/middleware"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

func createOrderHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.CreateOrderRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.Errorw(ctx, "error occured while decoding request",
				zap.Error(err),
			)
			middleware.ErrorResponse(ctx, w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody)
			return
		}

		err = req.Validate()
		if err != nil {
			logger.Errorw(ctx, "error occured while validating create order request",
				zap.Error(err),
			)

			if err == apperrors.ErrNoProductsToOrder {
				middleware.ErrorResponse(ctx, w, http.StatusBadRequest, err)
				return
			}

			middleware.ErrorResponse(ctx, w, http.StatusUnprocessableEntity, err)
			return
		}

		orderInfo, err := orderSvc.CreateOrder(ctx, req)
		if err != nil {
			logger.Errorw(ctx, "error occured while creating order",
				zap.Error(err),
			)
			statusCode, err := apperrors.MapError(err)
			middleware.ErrorResponse(ctx, w, statusCode, err)
			return
		}

		middleware.SuccessResponse(ctx, w, http.StatusCreated, orderInfo)
	}
}

func getOrderDetailsHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		rawOrderID := vars["id"]
		orderID, err := strconv.Atoi(rawOrderID)
		if err != nil {
			logger.Errorw(ctx, "error occured while converting orderID to an integer",
				zap.Error(err),
				zap.String("id", rawOrderID),
			)

			middleware.ErrorResponse(ctx, w, http.StatusBadRequest, apperrors.ErrInvalidRequestParam)
			return
		}

		response, err := orderSvc.GetOrderDetailsByID(ctx, int64(orderID))
		if err != nil {
			logger.Errorw(ctx, "error occured while fetching order info",
				zap.Error(err),
			)

			statusCode, errResponse := apperrors.MapError(err)
			middleware.ErrorResponse(ctx, w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(ctx, w, http.StatusOK, response)
	}
}

func updateOrderStatusHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.UpdateOrderStatusRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.Errorw(ctx, "error occured while decoding request",
				zap.Error(err),
			)
			middleware.ErrorResponse(ctx, w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody)
			return
		}

		err = req.Validate()
		if err != nil {
			logger.Errorw(ctx, "error occured validating update order request",
				zap.Error(err),
			)
			middleware.ErrorResponse(ctx, w, http.StatusBadRequest, err)
			return
		}

		orderInfo, err := orderSvc.UpdateOrderStatus(ctx, req.OrderID, req.Status)
		if err != nil {
			logger.Errorw(ctx, "error occured while updating order status",
				zap.Error(err),
			)
			statusCode, err := apperrors.MapError(err)
			middleware.ErrorResponse(ctx, w, statusCode, err)
			return
		}

		middleware.SuccessResponse(ctx, w, http.StatusOK, orderInfo)
	}
}

func listOrdersHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		response, err := orderSvc.ListOrders(ctx)
		if err != nil {
			logger.Errorw(ctx, "error occured while fetching orders list",
				zap.Error(err),
			)

			middleware.ErrorResponse(ctx, w, http.StatusInternalServerError, apperrors.ErrInternalServerError)
			return
		}

		middleware.SuccessResponse(ctx, w, http.StatusOK, response)
	}
}
