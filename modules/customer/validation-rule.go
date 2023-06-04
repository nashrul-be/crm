package customer

import (
	"errors"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type validateFunc func(entities.Customer, repositories.CustomerRepositoryInterface) (error, error)

func validateEmail(customer entities.Customer, customerRepo repositories.CustomerRepositoryInterface) (error, error) {
	exist, err := customerRepo.IsEmailExist(customer)
	if err != nil {
		return nil, err
	}
	if exist {
		return errors.New("email already taken"), nil
	}
	return nil, nil
}

func validateId(customer entities.Customer, customerRepo repositories.CustomerRepositoryInterface) (error, error) {
	exist, err := customerRepo.IsExist(customer.ID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return errors.New("customer ID doesn't exist"), nil
	}
	return nil, nil
}
