package persistence

import (
	"context"

	model "github.com/yohannesgossaye/internal/domain/model/customeraccount"
)

type CustomerRepositary interface {
	CreateCustomer(ctx context.Context, req *model.Customer) (model.Customer, error)
	LoginCustomer(ctx context.Context, email string, password string) (model.Customer, error)
	VerifyCustomerOtp(ctx context.Context, email string, otpCode string) (model.Customer, error)
	UpdateCustomer(ctx context.Context, customer *model.Customer) error
}
