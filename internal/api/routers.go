package api

import (
	"net/http"

	"github.com/divyatambat/FarmersBasket/internal/app"

	"github.com/gorilla/mux"
)

func NewRouter(deps app.Dependencies) *mux.Router {
	// Create the main router
	router := mux.NewRouter()

	// Order APIs
	orderRouter := router.PathPrefix("/orders").Subrouter()
	orderRouter.HandleFunc("/", createOrderHandler(deps.OrderService)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/", listOrdersHandler(deps.OrderService)).Methods(http.MethodGet)
	orderRouter.HandleFunc("/{id}", getOrderDetailsHandler(deps.OrderService)).Methods(http.MethodGet)
	orderRouter.HandleFunc("/{id}/status", updateOrderStatusHandler(deps.OrderService)).Methods(http.MethodPatch)

	// Product APIs
	productRouter := router.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("/{id}", getProductHandler(deps.ProductService)).Methods(http.MethodGet)
	productRouter.HandleFunc("/", listProductHandler(deps.ProductService)).Methods(http.MethodGet)

	// User APIs
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/{id}", getUserHandler(deps.UserService)).Methods(http.MethodGet)
	userRouter.HandleFunc("/", listUsersHandler(deps.UserService)).Methods(http.MethodGet)

	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// middleware logic here
		next.ServeHTTP(w, r)
	})
}
