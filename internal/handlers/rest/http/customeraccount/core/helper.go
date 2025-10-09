package core

import (
	dto "github.com/yohannesgossaye/internal/domain/dto/customeraccount"
	"github.com/yohannesgossaye/internal/domain/model/customeraccount"
)

func MapCreateCustomerRequestToModel(req dto.CreateCustomer) customeraccount.Customer {
	return customeraccount.Customer{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		Country:     req.Country,
		City:        req.City,
		Address:     req.Address,
		ZipCode:     req.ZipCode,
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
	}
}
