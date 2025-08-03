package service

import (
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
	"context"
)

type CustomerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

func (c CustomerService) Index(ctx context.Context)  ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil {
		// Handle error (e.g., log it, return an error response)
		return nil, err
	}

	var customerData []dto.CustomerData

	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}
	
	return customerData, nil
}