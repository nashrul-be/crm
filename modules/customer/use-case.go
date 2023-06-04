package customer

import (
	"errors"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type UseCaseInterface interface {
	ValidateCustomer(customer entities.Customer, validations ...validateFunc) (error, error)
	GetByID(id uint) (customer entities.Customer, err error)
	GetAll(limit, offset uint) (customers []entities.Customer, err error)
	GetAllByEmail(email string, limit, offset uint) (customers []entities.Customer, err error)
	GetAllByName(name string, limit, offset uint) (customers []entities.Customer, err error)
	CreateCustomer(customer *entities.Customer) (err error)
	UpdateCustomer(customer *entities.Customer) (err error)
	DeleteCustomer(id uint) (err error)
}

func NewUseCase(repositoryInterface repositories.CustomerRepositoryInterface) UseCaseInterface {
	return customerUseCase{customerRepository: repositoryInterface}
}

type customerUseCase struct {
	customerRepository repositories.CustomerRepositoryInterface
}

func (uc customerUseCase) ValidateCustomer(customer entities.Customer, validations ...validateFunc) (error, error) {
	for _, validation := range validations {
		validationError, err := validation(customer, uc.customerRepository)
		if err != nil {
			return nil, err
		}
		if validationError != nil {
			return validationError, nil
		}
	}
	return nil, nil
}

func (uc customerUseCase) GetByID(id uint) (customer entities.Customer, err error) {
	return uc.customerRepository.GetByID(id)
}

func (uc customerUseCase) GetAll(limit, offset uint) (customers []entities.Customer, err error) {
	return uc.customerRepository.GetAll(limit, offset)
}

func (uc customerUseCase) GetAllByEmail(email string, limit, offset uint) (customers []entities.Customer, err error) {
	return uc.customerRepository.GetAllByEmail(email, limit, offset)
}

func (uc customerUseCase) GetAllByName(name string, limit, offset uint) (customers []entities.Customer, err error) {
	return uc.customerRepository.GetAllByName(name, limit, offset)
}

func (uc customerUseCase) CreateCustomer(customer *entities.Customer) (err error) {
	return uc.customerRepository.Create(customer)
}

func (uc customerUseCase) UpdateCustomer(customer *entities.Customer) (err error) {
	exist, err := uc.customerRepository.IsExist(customer.ID)
	if err != nil {
		return
	}
	if !exist {
		return errors.New("customer doesn't exist")
	}
	err = uc.customerRepository.Save(customer)
	return
}

func (uc customerUseCase) DeleteCustomer(id uint) (err error) {
	err = uc.customerRepository.Delete(id)
	return
}
