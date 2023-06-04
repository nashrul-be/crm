package customer

import (
	"fmt"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"net/http"
)

type ControllerInterface interface {
	GetByID(id uint) (dto.BaseResponse, error)
	CreateCustomer(req CreateRequest) (dto.BaseResponse, error)
	GetAll(req PaginationRequest) (dto.BaseResponse, error)
	UpdateOrCreateCustomer(id uint, req CreateRequest) (dto.BaseResponse, error)
	DeleteCustomer(id uint) error
}

func NewController(useCaseInterface UseCaseInterface) ControllerInterface {
	return controller{
		customerUseCase: useCaseInterface,
	}
}

type controller struct {
	customerUseCase UseCaseInterface
}

func (c controller) GetByID(id uint) (dto.BaseResponse, error) {
	customer, err := c.customerUseCase.GetByID(id)
	if err != nil {
		return dto.ErrorNotFound(fmt.Sprintf("Customer %d doesn't exist", id)), err
	}
	customerRepresentation := mapCustomerToResponse(customer)
	return dto.Success("Success retrieve data", customerRepresentation), nil
}

func (c controller) GetAll(req PaginationRequest) (dto.BaseResponse, error) {
	var customers []entities.Customer
	var err error
	offset := (req.Page - 1) * req.PerPage
	switch {
	case req.Email != "":
		customers, err = c.customerUseCase.GetAllByEmail(req.Email+"%", req.PerPage, offset)
	case req.Name != "":
		customers, err = c.customerUseCase.GetAllByName(req.Name+"%", req.PerPage, offset)
	default:
		customers, err = c.customerUseCase.GetAllByEmail("%", req.PerPage, offset)
	}
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	result := make([]Representation, 0)
	for _, customer := range customers {
		result = append(result, mapCustomerToResponse(customer))
	}
	return dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success retrieve customers",
		Data:    result,
	}, nil
}

func (c controller) CreateCustomer(req CreateRequest) (dto.BaseResponse, error) {
	customer := mapCreateRequestToCustomer(req)
	exist, err := c.customerUseCase.IsEmailExist(customer)
	if err != nil {
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't create customer",
		}, err
	}
	if exist {
		return dto.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email already exist",
		}, nil
	}
	if err := c.customerUseCase.CreateCustomer(&customer); err != nil {
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't create customer",
		}, err
	}
	response := mapCustomerToResponse(customer)
	return dto.Created("Success create customer", response), nil
}

func (c controller) UpdateOrCreateCustomer(id uint, req CreateRequest) (dto.BaseResponse, error) {
	customer := mapCreateRequestToCustomer(req)
	customer.ID = id
	exist, err := c.customerUseCase.IsEmailExist(customer)
	if err != nil {
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't create customer",
		}, err
	}
	if exist {
		return dto.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email already exist",
		}, nil
	}
	if err := c.customerUseCase.UpdateOrCreateCustomer(&customer); err != nil {
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't update customer",
		}, err
	}
	response := mapCustomerToResponse(customer)
	return dto.Success("Success update customer", response), nil
}

func (c controller) DeleteCustomer(id uint) error {
	if err := c.customerUseCase.DeleteCustomer(id); err != nil {
		return err
	}
	return nil
}
