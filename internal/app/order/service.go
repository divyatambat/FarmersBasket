package order

import (
	"context"
	"fmt"
	"time"

	"github.com/divyatambat/FarmersBasket/internal/app/product"
	"github.com/divyatambat/FarmersBasket/internal/pkg/apperrors"
	"github.com/divyatambat/FarmersBasket/internal/pkg/dto"
	"github.com/divyatambat/FarmersBasket/internal/repository"
)

var now = time.Now

type service struct {
	orderRepo      repository.OrderStorer
	orderItemsRepo repository.OrderItemStorer
	productSvc     product.Service
}

type Service interface {
	CreateOrder(ctx context.Context, orderDetails dto.CreateOrderRequest) (dto.Order, error)
	GetOrderDetailsByID(ctx context.Context, orderID int64) (dto.Order, error)
	ListOrders(ctx context.Context) ([]dto.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID int64, status string) (dto.Order, error)
	UpdatePaymentStatus(ctx context.Context, orderID int64, status string) (dto.Order, error)
}

func NewService(orderRepo repository.OrderStorer, orderItemsRepo repository.OrderItemStorer,
	productSvc product.Service) Service {
	return &service{
		orderRepo:      orderRepo,
		orderItemsRepo: orderItemsRepo,
		productSvc:     productSvc,
	}
}

func (os *service) CreateOrder(ctx context.Context, orderDetails dto.CreateOrderRequest) (order dto.Order, err error) {
	// Calculate order value directly within the function
	var orderValue float64
	productQuantityUpdates := make([]product.UpdateProductQuantity, 0)
	for _, p := range orderDetails.Product {
		productDetails, err := os.productSvc.GetProductByID(ctx, p.ID)
		if err != nil {
			return dto.Order{}, fmt.Errorf("error getting product details: %w", err)
		}

		itemValue := productDetails.Price * float64(p.Quantity)
		orderValue += itemValue

		productQuantityUpdates = append(productQuantityUpdates, product.UpdateProductQuantity{
			ProductID: p.ID,
			Quantity:  p.Quantity,
		})
	}

	// Create order object with calculated value
	orderRepoObj := repository.Order{
		Total_Amount: orderValue,
		ID:           uint(orderValue),
		Order_Status: ListOrderStatus[OrderPlaced],
	}

	// Set Order Status to Placed
	orderRepoObj.Order_Status = ListOrderStatus[OrderPlaced]

	// Inserting Order in Database
	orderDB, err := os.orderRepo.CreateOrder(ctx, orderRepoObj)
	if err != nil {
		return dto.Order{}, err
	}

	orderItems := make([]repository.OrderItem, 0)

	updatedProductInfo := []product.UpdateProductQuantity{}
	for _, item := range orderItems {
		// Only update quantities if order is not cancelled or returned
		if MapOrderStatus[order.OrderStatus] != OrderCancelled && MapOrderStatus[order.OrderStatus] != OrderReturned {
			updatedProductInfo = append(updatedProductInfo, product.UpdateProductQuantity{
				ProductID: item.Product_ID,
				Quantity:  -item.Quantity, // Decrease quantity for purchase
			})
		}
	}

	// Update Product quantity in database
	productQuantityMap := make(map[int64]int64)
	for _, p := range updatedProductInfo {
		productQuantityMap[p.ProductID] = p.Quantity
	}

	err = os.productSvc.UpdateProductQuantity(ctx, productQuantityMap)
	if err != nil {
		return dto.Order{}, err
	}

	for _, p := range orderDetails.Product {
		orderItem := repository.OrderItem{
			ID:          0,
			Customer_ID: 0,
			Product_ID:  p.ID,
			Quantity:    p.Quantity,
			Order_Date:  time.Time{},
			Order_type:  "",
		}
		orderItems = append(orderItems, orderItem)
	}

	// Inserting order items in database
	err = os.orderItemsRepo.StoreOrderItems(ctx, orderItems)
	if err != nil {
		return dto.Order{}, err
	}

	order = MapOrderRepoToOrderDto(orderDB, orderItems...)
	return order, nil
}

func (os *service) GetOrderDetailsByID(ctx context.Context, orderID int64) (order dto.Order, err error) {
	orderInfoDB, err := os.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return dto.Order{}, err
	}

	if orderInfoDB.ID == 0 {
		return dto.Order{}, apperrors.OrderNotFound{ID: orderID}
	}

	orderItemsDB, err := os.orderItemsRepo.GetOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		return dto.Order{}, err
	}

	order = MapOrderRepoToOrderDto(orderInfoDB, orderItemsDB...) //checking for errors in mapping
	return order, nil
}

func (os *service) ListOrders(ctx context.Context) ([]dto.Order, error) {
	orderList := make([]dto.Order, 0)

	orderListDB, err := os.orderRepo.ListOrders(ctx)
	if err != nil {
		return orderList, err
	}

	for _, order := range orderListDB {
		orderList = append(orderList, MapOrderRepoToOrderDto(order))
	}

	return orderList, nil
}

func (os *service) UpdateOrderStatus(ctx context.Context, orderID int64, status string) (order dto.Order, err error) {
	// get order transaction
	orderInfoDB, err := os.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return dto.Order{}, err
	}

	// Order status validation
	if !validateUpdateOrderStatusRequest(status, orderInfoDB.Order_Status) {
		return dto.Order{}, apperrors.OrderUpdationInvalid{
			ID:             orderID,
			CurrentState:   orderInfoDB.Order_Status,
			RequestedState: status,
		}
	}

	// Update order status in database
	err = os.orderRepo.UpdateOrderStatus(ctx, orderID, status)
	if err != nil {
		return dto.Order{}, fmt.Errorf("error occurred while updating order status: %w", err)
	}

	// Update product quantity if order cancelled or returned
	if MapOrderStatus[status] == OrderCancelled || MapOrderStatus[status] == OrderReturned {
		_, err := os.orderItemsRepo.GetOrderItemsByOrderID(ctx, orderID)
		if err != nil {
			return dto.Order{}, fmt.Errorf("error occurred while fetching order items: %w", err)
		}
	}

	order = MapOrderRepoToOrderDto(orderInfoDB)
	return order, nil
}

func (os *service) UpdatePaymentStatus(ctx context.Context, orderID int64, status string) (order dto.Order, err error) {
	// Implement payment status update logic here
	return dto.Order{}, nil
}
