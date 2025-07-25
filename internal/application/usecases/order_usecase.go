package usecases

import (
	"errors"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/repositories"
)

type OrderUseCase interface {
	CreateOrder(request *dto.CreateOrderRequest) (*dto.OrderResponse, error)
	GetOrderByID(id uint64) (*dto.OrderResponse, error)
	GetOrdersByCPF(cpf string) ([]*dto.OrderResponse, error)
	GetOrdersByCustomerID(customerID uint64) ([]*dto.OrderResponse, error)
	GetAllOrders() ([]*dto.OrderResponse, error)
	GetOrdersForKitchen() ([]*dto.OrderResponse, error)
	UpdateOrderStatus(id uint64, request *dto.UpdateOrderStatusRequest) (*dto.OrderResponse, error)
	DeleteOrder(id uint64) error
}

type orderUseCase struct {
	orderRepo     repositories.OrderRepository
	orderItemRepo repositories.OrderItemRepository
	productRepo   repositories.ProductRepository
}

func NewOrderUseCase(
	orderRepo repositories.OrderRepository,
	orderItemRepo repositories.OrderItemRepository,
	productRepo repositories.ProductRepository,
) OrderUseCase {
	return &orderUseCase{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		productRepo:   productRepo,
	}
}

func (uc *orderUseCase) CreateOrder(request *dto.CreateOrderRequest) (*dto.OrderResponse, error) {
	order := entities.NewOrder(request.CustomerId, request.CPF)

	var totalPrice float32
	for _, itemReq := range request.Items {
		product, err := uc.productRepo.GetByID(itemReq.ProductID)
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

	err := uc.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	for _, item := range order.Items {
		item.OrderID = order.ID
		err := uc.orderItemRepo.Create(&item)
		if err != nil {
			return nil, err
		}
	}

	return uc.buildOrderResponse(order), nil
}

func (uc *orderUseCase) GetOrderByID(id uint64) (*dto.OrderResponse, error) {
	order, err := uc.orderRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	items, err := uc.orderItemRepo.GetByOrderID(id)
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
	orders, err := uc.orderRepo.GetByCPF(cpf)
	if err != nil {
		return nil, err
	}

	var response []*dto.OrderResponse
	for _, order := range orders {
		items, err := uc.orderItemRepo.GetByOrderID(order.ID)
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
	orders, err := uc.orderRepo.GetByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	var response []*dto.OrderResponse
	for _, order := range orders {
		items, err := uc.orderItemRepo.GetByOrderID(order.ID)
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
	orders, err := uc.orderRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var response []*dto.OrderResponse
	for _, order := range orders {
		items, err := uc.orderItemRepo.GetByOrderID(order.ID)
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
	orders, err := uc.orderRepo.GetPendingOrdersForKitchen()
	if err != nil {
		return nil, err
	}

	var response []*dto.OrderResponse
	for _, order := range orders {
		items, err := uc.orderItemRepo.GetByOrderID(order.ID)
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
	order, err := uc.orderRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	order.UpdateStatus(entities.OrderStatus(request.Status))

	err = uc.orderRepo.Update(order)
	if err != nil {
		return nil, err
	}

	items, err := uc.orderItemRepo.GetByOrderID(id)
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
	order, err := uc.orderRepo.GetByID(id)
	if err != nil {
		return err
	}

	if order == nil {
		return errors.New("order not found")
	}

	return uc.orderRepo.Delete(id)
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
