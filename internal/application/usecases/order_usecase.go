package usecases

import (
	"errors"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output"
)

type orderUseCase struct {
	orderGateway     output.OrderGateway
	orderItemGateway output.OrderItemGateway
	productGateway   output.ProductGateway
}

func NewOrderUseCase(
	orderGateway output.OrderGateway,
	orderItemGateway output.OrderItemGateway,
	productGateway output.ProductGateway,
) input.OrderUseCase {
	return &orderUseCase{
		orderGateway:     orderGateway,
		orderItemGateway: orderItemGateway,
		productGateway:   productGateway,
	}
}

func (uc *orderUseCase) CreateOrder(request *dto.CreateOrderRequest) (*dto.OrderResponse, error) {
	order := entities.NewOrder(request.CustomerId, request.CPF)

	var totalPrice float32
	for _, itemReq := range request.Items {
		product, err := uc.productGateway.GetByID(itemReq.ProductID)
		if err != nil || product == nil {
			return nil, errors.New("product not found")
		}

		orderItem := entities.NewOrderItem(order.ID, itemReq.ProductID, itemReq.Quantity, product.Price)

		if !orderItem.IsValid() {
			return nil, errors.New("invalid order item data")
		}

		order.AddItem(*orderItem)
		totalPrice += orderItem.CalculateSubtotal()
	}

	if !order.IsValid() {
		return nil, errors.New("invalid order data")
	}

	err := uc.orderGateway.Create(order)
	if err != nil {
		return nil, err
	}

	for _, item := range order.Items {
		item.OrderID = order.ID
		err := uc.orderItemGateway.Create(&item)
		if err != nil {
			return nil, err
		}
	}

	return uc.buildOrderResponse(order), nil
}

func (uc *orderUseCase) GetOrderByID(id uint64) (*dto.OrderResponse, error) {
	order, err := uc.orderGateway.GetByID(id)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	items, err := uc.orderItemGateway.GetByOrderID(id)
	if err != nil {
		return nil, err
	}

	var orderItems []entities.OrderItem
	for _, item := range items {
		orderItems = append(orderItems, *item)
	}
	order.Items = orderItems

	return uc.buildOrderResponse(order), nil
}

func (uc *orderUseCase) GetOrdersByCPF(cpf string) ([]*dto.OrderResponse, error) {
	orders, err := uc.orderGateway.GetByCPF(cpf)
	if err != nil {
		return nil, err
	}

	var response []*dto.OrderResponse
	for _, order := range orders {
		items, err := uc.orderItemGateway.GetByOrderID(order.ID)
		if err != nil {
			continue
		}

		var orderItems []entities.OrderItem
		for _, item := range items {
			orderItems = append(orderItems, *item)
		}
		order.Items = orderItems

		response = append(response, uc.buildOrderResponse(order))
	}

	return response, nil
}

func (uc *orderUseCase) GetOrdersByCustomerID(customerID uint64) ([]*dto.OrderResponse, error) {
	orders, err := uc.orderGateway.GetByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	var response []*dto.OrderResponse
	for _, order := range orders {
		items, err := uc.orderItemGateway.GetByOrderID(order.ID)
		if err != nil {
			continue
		}

		var orderItems []entities.OrderItem
		for _, item := range items {
			orderItems = append(orderItems, *item)
		}
		order.Items = orderItems

		response = append(response, uc.buildOrderResponse(order))
	}

	return response, nil
}

func (uc *orderUseCase) GetAllOrders() ([]*dto.OrderResponse, error) {
	orders, err := uc.orderGateway.GetAll()
	if err != nil {
		return nil, err
	}

	var response []*dto.OrderResponse
	for _, order := range orders {
		items, err := uc.orderItemGateway.GetByOrderID(order.ID)
		if err != nil {
			continue
		}

		var orderItems []entities.OrderItem
		for _, item := range items {
			orderItems = append(orderItems, *item)
		}
		order.Items = orderItems

		response = append(response, uc.buildOrderResponse(order))
	}

	return response, nil
}

func (uc *orderUseCase) GetOrdersForKitchen() ([]*dto.OrderResponse, error) {
	orders, err := uc.orderGateway.GetPendingOrdersForKitchen()
	if err != nil {
		return nil, err
	}

	var response []*dto.OrderResponse
	for _, order := range orders {
		items, err := uc.orderItemGateway.GetByOrderID(order.ID)
		if err != nil {
			continue
		}

		var orderItems []entities.OrderItem
		for _, item := range items {
			orderItems = append(orderItems, *item)
		}
		order.Items = orderItems

		response = append(response, uc.buildOrderResponse(order))
	}

	return response, nil
}

func (uc *orderUseCase) UpdateOrderStatus(id uint64, request *dto.UpdateOrderStatusRequest) (*dto.OrderResponse, error) {
	order, err := uc.orderGateway.GetByID(id)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	order.UpdateStatus(entities.OrderStatus(request.Status))

	err = uc.orderGateway.Update(order)
	if err != nil {
		return nil, err
	}

	items, err := uc.orderItemGateway.GetByOrderID(id)
	if err != nil {
		return nil, err
	}

	var orderItems []entities.OrderItem
	for _, item := range items {
		orderItems = append(orderItems, *item)
	}
	order.Items = orderItems

	return uc.buildOrderResponse(order), nil
}

func (uc *orderUseCase) DeleteOrder(id uint64) error {
	order, err := uc.orderGateway.GetByID(id)
	if err != nil {
		return err
	}

	if order == nil {
		return errors.New("order not found")
	}

	return uc.orderGateway.Delete(id)
}

func (uc *orderUseCase) buildOrderResponse(order *entities.Order) *dto.OrderResponse {
	var items []dto.OrderItemResponse
	for _, item := range order.Items {
		items = append(items, dto.OrderItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Subtotal:  item.CalculateSubtotal(),
		})
	}

	return &dto.OrderResponse{
		ID:         order.ID,
		CustomerId: order.CustomerId,
		CPF:        order.CPF,
		Status:     string(order.Status),
		Items:      items,
		Total:      order.CalculateTotal(),
		CreatedAt:  order.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:  order.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
