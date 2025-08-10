package service

import (
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
	"context"
	"github.com/google/uuid"
	"database/sql"
	"time"
	"errors"
    // "fmt"
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

func (c CustomerService) Create(ctx context.Context, req dto.CreateCustomerRequest) ([]dto.CustomerData, error) {
    customer := domain.Customer{
        ID:        uuid.New().String(),
        Code:      req.Code,
        Name:      req.Name,
        CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
    }

    err := c.customerRepository.Save(ctx, &customer)
    if err != nil {
        return nil, err
    }

    // Misal ingin mengembalikan customer yang baru dibuat dalam slice:
    createdCustomer := dto.CustomerData{
        ID:   customer.ID,
        Code: customer.Code,
        Name: customer.Name,
        // tambahkan field sesuai struct CustomerData
    }
    return []dto.CustomerData{createdCustomer}, nil
}

func (c CustomerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) ([]dto.CustomerData, error) {
    // Cari data customer
    exist, err := c.customerRepository.FindById(ctx, req.ID)
    // fmt.Println(err)

    // Jika customer tidak ditemukan
    if err != nil && exist.ID == "" {
        return nil, errors.New("Data customer tidak ditemukan!.")
    }
    
    if err != nil {
        return nil, err
    }

    // Update data sesuai request
    exist.Code = req.Code
    exist.Name = req.Name
    exist.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

    // Simpan perubahan
    err = c.customerRepository.Update(ctx, &exist)

    if err != nil {
        return nil, err
    }

    // Buat response DTO
    updatedCustomer := dto.CustomerData{
        ID:   exist.ID,
        Code: exist.Code,
        Name: exist.Name,
        // tambahkan field lain sesuai kebutuhan
    }

    return []dto.CustomerData{updatedCustomer}, nil
}

func (c CustomerService) Delete (ctx context.Context, id string) error {
    exist, err := c.customerRepository.FindById(ctx,id)

    if err != nil && exist.ID == "" {
        return  errors.New("Data customer tidak ditemukan!.")
    }
    
    if err != nil {
        return err
    }

    return c.customerRepository.Delete(ctx,id)
}

func (c CustomerService) Show (ctx context.Context, id string) (dto.CustomerData, error) {
    exist, err := c.customerRepository.FindById(ctx,id)

    if err != nil && exist.ID == "" {
        return dto.CustomerData{}, errors.New("Data customer tidak ditemukan!.")
    }
    
    if err != nil {
        return dto.CustomerData{}, err
    }

    return dto.CustomerData{
        ID:exist.ID,
        Code:exist.Code,
        Name:exist.Name,
    }, nil
}