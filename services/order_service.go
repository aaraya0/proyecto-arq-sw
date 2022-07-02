package services

import (
	"time"

	orderDetailCliente "github.com/aaraya0/arq-software/Integrador1/clients/detail"
	orderCliente "github.com/aaraya0/arq-software/Integrador1/clients/order"
	productCliente "github.com/aaraya0/arq-software/Integrador1/clients/product"
	"github.com/aaraya0/arq-software/Integrador1/dto"
	"github.com/aaraya0/arq-software/Integrador1/model"
	e "github.com/aaraya0/arq-software/Integrador1/utils/errors"
	//"github.com/dgrijalva/jwt-go"
)

type orderService struct{}

type orderServiceInterface interface {
	InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError)
	GetOrdersByUserId(token int) (dto.OrdersDto, e.ApiError) //token cambiado de string a int temporalmente. revisar
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError) {

	var order model.Order

	order.Date = time.Now()

	order.User_Id = orderDto.User_Id

	order = orderCliente.InsertOrder(order)

	orderDto.Id = order.Id

	var orderDetails model.Details
	var montofinal float32
	for _, orderDetailDto := range orderDto.OrderDetails {
		var orderDetail model.Detail

		orderDetail.Product_Id = orderDetailDto.Product_Id

		var product model.Product = productCliente.GetProductById(orderDetail.Product_Id)
		orderDetail.Name = product.Title
		orderDetail.UnitPrice = product.Price

		orderDetail.Quantity = orderDetailDto.Quantity
		orderDetail.Total = orderDetail.UnitPrice * orderDetail.Quantity

		orderDetail.Order_Id = orderDto.Id
		for i := 0; i < len(orderDto.OrderDetails); i++ {
			detail := orderDto.OrderDetails[i]
			product := productCliente.GetProductById(detail.Product_Id)
			if float32(product.Stock) < detail.Quantity {
				orderDto.Id = 0
				return orderDto, e.NewConflictApiError("Not enough stock on product: " + product.Title)
			}

			montofinal += orderDetail.Total

		}

		orderDetails = append(orderDetails, orderDetail)
	}

	order.TotalAmount = orderCliente.UpdateMontoFinal(montofinal, order.Id)

	orderDetails = orderDetailCliente.InsertOrderDetails(orderDetails)

	var orderResponseDto dto.OrderDto

	orderResponseDto.Date = order.Date
	orderResponseDto.Id = order.Id
	orderResponseDto.User_Id = order.User_Id
	orderResponseDto.TotalAmount = order.TotalAmount

	for _, orderDetail := range orderDetails {
		var orderDetailDto dto.DetailDto

		orderDetailDto.Id = orderDetail.Id
		orderDetailDto.Quantity = orderDetail.Quantity
		orderDetailDto.Product_Id = orderDetail.Product_Id
		orderDetailDto.UnitPrice = orderDetail.UnitPrice
		orderDetailDto.Total = orderDetail.Total
		orderDetailDto.Order_Id = orderDetail.Order_Id
		orderDetailDto.Name = orderDetail.Name

		orderResponseDto.OrderDetails = append(orderResponseDto.OrderDetails, orderDetailDto)
	}

	return orderResponseDto, nil
}

//Buscar orden por IDuser

func (s *orderService) GetOrdersByUserId(token int) (dto.OrdersDto, e.ApiError) { //token cambiado de string a int temporalmente (login sin token)
	//var userId float64
	/*tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, e.NewUnauthorizedApiError("Unauthorized")
		}
		return nil, e.NewUnauthorizedApiError("Unauthorized")
	}

	if !tkn.Valid {
		return nil, e.NewUnauthorizedApiError("Unauthorized")

	}
	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {

		userId = (claims["user_id"].(float64))

	} else {
		return nil, e.NewUnauthorizedApiError("Unauthorized")
	}
	*/
	//var IdUserX int = int(userId)   //temporalmente comentado porq el login es sin token. revisar
	var IdUserX int = token
	var orders model.Orders = orderCliente.GetOrdersByUserId(IdUserX)
	var ordersDto dto.OrdersDto

	if len(orders) == 0 {
		return ordersDto, e.NewBadRequestApiError("order not found")
	}

	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id = order.Id
		orderDto.Date = order.Date
		orderDto.TotalAmount = order.TotalAmount

		var orderDetails model.Details = orderDetailCliente.GetOrderDetailByOrderId(order.Id)
		for _, orderDetail := range orderDetails {
			var orderDetailDto dto.DetailDto

			orderDetailDto.Id = orderDetail.Id
			orderDetailDto.Quantity = orderDetail.Quantity
			orderDetailDto.Product_Id = orderDetail.Product_Id
			orderDetailDto.UnitPrice = orderDetail.UnitPrice
			orderDetailDto.Total = orderDetail.Total
			orderDetailDto.Order_Id = orderDetail.Order_Id
			orderDetailDto.Name = orderDetail.Name

			orderDto.OrderDetails = append(orderDto.OrderDetails, orderDetailDto)
		}
		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}
