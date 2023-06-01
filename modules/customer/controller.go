package customer

import (
	"fmt"
	"nashrul-be/crm/dto"
	"net/http"
)

type ControllerInterface interface {
	GetByID(id uint) (any, error)
	CreateCustomer(req CreateRequest) (any, error)
	UpdateOrCreateCustomer(id uint, req CreateRequest) (any, error)
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

func (c controller) GetByID(id uint) (any, error) {
	customer, err := c.customerUseCase.GetByID(id)
	if err != nil {
		return dto.ErrorNotFound(fmt.Sprintf("Customer %d doesn't exist", id)), err
	}
	return dto.Success("Success retrieve data", customer), nil
}

func (c controller) CreateCustomer(req CreateRequest) (any, error) {
	customer := mapCreateRequestToCustomer(req)
	if err := c.customerUseCase.CreateCustomer(&customer); err != nil {
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't create customer",
		}, err
	}
	response := mapCustomerToResponse(customer)
	return dto.Created("Success create customer", response), nil
}

func (c controller) UpdateOrCreateCustomer(id uint, req CreateRequest) (any, error) {
	customer := mapCreateRequestToCustomer(req)
	customer.ID = id
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
