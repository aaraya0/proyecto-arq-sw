package services

/*import (
	//falta client
	//ordCliente "github.com/aaraya0/arq-software/Integrador1/clients/order"
	//userCliente 	"github.com/aaraya0/arq-software/Integrador1/clients/user"
	//"github.com/aaraya0/arq-software/Integrador1/dto"
	//"github.com/aaraya0/arq-software/Integrador1/model"
	e "github.com/aaraya0/arq-software/Integrador1/utils/errors"
)

type orderService struct{}

type orderServiceInterface interface {
	GetOrdersByUName(id string) (dto.OrdersDto, e.ApiError)
	AddOrder (orderDto dto.OrderDto) (dto.OrderDtoResp, e.ApiError)
	//InsertProduct
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}
/*func (p *orderService) GetOrderById(id int) (dto.OrderDto, e.ApiError) {
	var order model.Order = ordCliente.GetOrderById(id)

	var orderDto dto.OrderDto
	var userDto  dto.UserDto
	if order.Id == 0 {
		return orderDto, e.NewBadRequestApiError("order not found")
	}

	orderDto.Id = order.Id

	orderDto.Total = order.Total
	//orderDto.OrderDetail = order.OrderDetail
	return orderDto, nil
}*/
/*func (s *orderService) AddOrder(orderDto dto.OrderDto) (dto.OrderDtoResp, e.ApiError)  {

	var order model.Order
	var orderDtoResp dto.OrderDtoResp
	var total float32
	total = 0
	order.User_id= orderDto.User

	....

}
*/
