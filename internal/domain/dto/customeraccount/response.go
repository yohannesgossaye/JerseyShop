package customeraccount

import (
	model "github.com/yohannesgossaye/internal/domain/model/customeraccount"
)

type CustomerResponse struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	IsVerified  bool   `json:"is_verified"`
}

type LoginResponse struct {
	AccessToken string           `json:"access_token"`
	Customer    CustomerResponse `json:"customer"`
}
type VerifyOtpResponse struct {
	AccessToken string         `json:"access_token"`
	Customer    model.Customer `json:"customer"`
}
