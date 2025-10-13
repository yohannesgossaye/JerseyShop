package temporalworkflow

import (
	"github.com/yohannesgossaye/internal/domain/model/customeraccount"
)

// map workflow to model
func MapWorkflowToModel(input UserWorkflowInput) customeraccount.Customer {
	return customeraccount.Customer{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		DateOfBirth: input.DateOfBirth,
		Country:     input.Country,
		City:        input.City,
		Address:     input.Address,
		ZipCode:     input.ZipCode,
		Gender:      input.Gender,
		OtpCode:     input.OtpCode,
	}
}
