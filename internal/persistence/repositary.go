package persistence

import (
	"context"

	modeladmin "github.com/yohannesgossaye/internal/domain/model/admin"
	model "github.com/yohannesgossaye/internal/domain/model/customeraccount"
)

type CustomerRepositary interface {
	CreateCustomer(ctx context.Context, req *model.Customer) (model.Customer, error)
	LoginCustomer(ctx context.Context, email string, password string) (model.Customer, error)
	VerifyCustomerOtp(ctx context.Context, email string, otpCode string) (model.Customer, error)
	UpdateCustomer(ctx context.Context, customer *model.Customer) error
}

type AdminRepositary interface {
	CreateAdmin(ctx context.Context, req *modeladmin.Admin) (modeladmin.Admin, error)
	GetAdmin(ctx context.Context, id string) (modeladmin.Admin, error)
	GetAdminAccounts(ctx context.Context) ([]modeladmin.Admin, error)
	LoginAdminAccount(ctx context.Context, email string, password string) (modeladmin.Admin, error)
	DeleteAdminAccount(ctx context.Context, id string) (string, error)
}
