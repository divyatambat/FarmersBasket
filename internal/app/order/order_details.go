package order

import (
	"github.com/divyatambat/FarmersBasket/internal/pkg/dto"
	"github.com/divyatambat/FarmersBasket/internal/repository"
)

type OrderStatus int

const (
	OrderCancelled OrderStatus = iota
	OrderPlaced
	OrderDispatched
	OrderCompleted
	OrderReturned
)

var MapOrderStatus = map[string]OrderStatus{
	"Cancelled":  OrderCancelled,
	"Placed":     OrderPlaced,
	"Dispatched": OrderDispatched,
	"Completed":  OrderCompleted,
	"Returned":   OrderReturned,
}

// the order of this slice needs to match
var ListOrderStatus = []string{
	"Cancelled",
	"Placed",
	"Dispatched",
	"Completed",
	"Returned",
}

func validateUpdateOrderStatusRequest(RequestOrderStatus, DBOrderStatus string) (isUpdateValid bool) {
	requestedOrderState := MapOrderStatus[RequestOrderStatus]
	currentOrderState := MapOrderStatus[DBOrderStatus]

	//do not update if order is already cancelled
	if currentOrderState == OrderCancelled {
		return false
	}

	//allow cancel only before order is completed
	if requestedOrderState == OrderCancelled && currentOrderState < OrderCompleted {
		return true
	}

	//donot update if requested state is same or lower to current state
	if currentOrderState >= requestedOrderState {
		return false
	}

	//order status update can only go one step forward
	if requestedOrderState != (currentOrderState + 1) {
		return false
	}

	return true
}

func MapOrderRepoToOrderDto(order repository.Order, orderItems ...repository.OrderItem) dto.Order {

	orderInfo := make([]dto.ProductInfo, 0)
	for _, orderItem := range orderItems {
		orderInfo = append(orderInfo, dto.ProductInfo{
			ID:       orderItem.Product_ID,
			Quantity: orderItem.Quantity,
		})
	}

	return dto.Order{
		ID:              int64(order.ID),
		TotalAmount:     order.Total_Amount,
		OrderStatus:     order.Order_Status,
		CustomerID:      order.Customer_ID,
		ProductID:       order.Product_ID,
		OrderDate:       order.Order_Date,
		PaymentOption:   order.Payment_option,
		PaymentStatus:   order.Payment_status,
		DeliveryAddress: order.Delivery_Address,
		OrderType:       order.Order_type,
		Quantity:        int64(order.Quantity),
	}
}
