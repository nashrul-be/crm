package customer

import "nashrul-be/crm/entities"

func mapCreateRequestToCustomer(request CreateRequest) entities.Customer {
	return entities.Customer{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Avatar:    request.Avatar,
	}
}

func mapUpdateRequestToCustomer(request UpdateRequest) entities.Customer {
	return entities.Customer{
		ID:        request.ID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Avatar:    request.Avatar,
	}
}

func mapCustomerToResponse(customer entities.Customer) Representation {
	return Representation{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}
}
