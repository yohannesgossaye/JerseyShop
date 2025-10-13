package core

import (
	"github.com/yohannesgossaye/internal/domain/model/customeraccount"
	workflows "github.com/yohannesgossaye/pkgs/temporalworkflow"
)

func Toworkflowinput(user customeraccount.Customer) workflows.UserWorkflowInput {
	return workflows.UserWorkflowInput{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		DateOfBirth: user.DateOfBirth,
		Country:     user.Country,
		City:        user.City,
		Address:     user.Address,
		ZipCode:     user.ZipCode,
		Gender:      user.Gender,
		OtpCode:     user.OtpCode,
	}
}
