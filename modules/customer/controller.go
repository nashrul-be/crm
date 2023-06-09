package customer

import (
	"encoding/json"
	"io"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"net/http"
)

type ControllerInterface interface {
	GetByID(id uint) (dto.BaseResponse, error)
	CreateCustomer(req CreateRequest) (dto.BaseResponse, error)
	GetAll(req PaginationRequest) (dto.BaseResponse, error)
	UpdateCustomer(req UpdateRequest) (dto.BaseResponse, error)
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
		return customerNotFound(), err
	}
	customerRepresentation := mapCustomerToResponse(customer)
	return dto.Success("Success retrieve data", customerRepresentation), nil
}

func (c controller) GetAll(req PaginationRequest) (dto.BaseResponse, error) {
	var customers []entities.Customer
	var err error
	response, err := http.Get("https://reqres.in/api/users?page=2")
	if err == nil {
		var jsonResponse ThirdPartyJSON
		body, err := io.ReadAll(response.Body)
		defer response.Body.Close()
		if err == nil {
			err = json.Unmarshal(body, &jsonResponse)
			if err == nil {
				for _, customer := range jsonResponse.Data {
					c.CreateCustomer(customer)
				}
			}
		}
	}
	offset := (req.Page - 1) * req.PerPage
	switch {
	case req.Email != "":
		customers, err = c.customerUseCase.GetAllByEmail(req.Email+"%", req.PerPage, offset)
	case req.Name != "":
		customers, err = c.customerUseCase.GetAllByName(req.Name+"%", req.PerPage, offset)
	default:
		customers, err = c.customerUseCase.GetAll(req.PerPage, offset)
	}
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	result := make([]Representation, 0)
	for _, customer := range customers {
		result = append(result, mapCustomerToResponse(customer))
	}
	return dto.Success("Success retrieve customers", customers), nil
}

func (c controller) CreateCustomer(req CreateRequest) (dto.BaseResponse, error) {
	customer := mapCreateRequestToCustomer(req)
	validationError, err := c.customerUseCase.ValidateCustomer(customer, validateEmail)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	if validationError != nil {
		return dto.ErrorBadRequest(validationError.Error()), nil
	}
	if err := c.customerUseCase.CreateCustomer(&customer); err != nil {
		return dto.ErrorInternalServerError(), err
	}
	response := mapCustomerToResponse(customer)
	return dto.Created("Success create customer", response), nil
}

func (c controller) UpdateCustomer(req UpdateRequest) (dto.BaseResponse, error) {
	customer := mapUpdateRequestToCustomer(req)
	validationError, err := c.customerUseCase.ValidateCustomer(customer, validateId, validateEmail)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	if validationError != nil {
		return dto.ErrorBadRequest(validationError.Error()), nil
	}
	if err := c.customerUseCase.UpdateCustomer(&customer); err != nil {
		return dto.ErrorInternalServerError(), err
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
