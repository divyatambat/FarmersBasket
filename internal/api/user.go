package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/divyatambat/FarmersBasket/internal/app/user"
	"github.com/divyatambat/FarmersBasket/internal/pkg/apperrors"
	"github.com/divyatambat/FarmersBasket/internal/pkg/logger"
	"github.com/divyatambat/FarmersBasket/internal/pkg/middleware"
	"go.uber.org/zap"
)

func getUserHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		rawUserID := vars["id"]

		userID, err := strconv.Atoi(rawUserID)
		if err != nil {
			logger.Errorw(ctx, "error occured while converting userID to an integer",
				zap.Error(err),
				zap.String("id", rawUserID),
			)

			middleware.ErrorResponse(ctx, w, http.StatusBadRequest, apperrors.ErrInvalidRequestParam)
			return
		}

		response, err := userSvc.GetUserByID(ctx, nil, uint(userID))
		if err != nil {
			logger.Errorw(ctx, "error occured while fetching user info",
				zap.Error(err),
			)

			statusCode, errResponse := apperrors.MapError(err)
			middleware.ErrorResponse(ctx, w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(ctx, w, http.StatusOK, response)
	}
}

func listUsersHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		response, err := userSvc.ListUsers(ctx)
		if err != nil {
			logger.Errorw(ctx, "error occured while fetching user list",
				zap.Error(err),
			)

			middleware.ErrorResponse(ctx, w, http.StatusInternalServerError, apperrors.ErrInternalServerError)
			return
		}

		middleware.SuccessResponse(ctx, w, http.StatusOK, response)
	}
}

// func RegisterUserHandlers(router *mux.Router, userSvc user.Service) {
// 	// routes using Gorilla Mux syntax
// 	router.HandleFunc("/users/{id}", getUserHandler(userSvc))
// 	router.HandleFunc("/users", listUsersHandler(userSvc))
// }
