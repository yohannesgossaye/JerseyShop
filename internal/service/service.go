package service

import (
	"context"

	custmeraccountdto "github.com/yohannesgossaye/internal/domain/dto/customeraccount"
	modeladmin "github.com/yohannesgossaye/internal/domain/model/admin"
	modelcustomer "github.com/yohannesgossaye/internal/domain/model/customeraccount"
)

type CustomerAccountService interface {
	CreateCustomer(ctx context.Context, req modelcustomer.Customer) (modelcustomer.Customer, error)
	LoginCustomer(ctx context.Context, email string, password string) (modelcustomer.Customer, error)
	VerifyCustomerOtp(ctx context.Context, email, otpCode string) (custmeraccountdto.VerifyOtpResponse, error)
	
}

type AdminAccountService interface {
	CreateAdmin(ctx context.Context, req *modeladmin.Admin) (modeladmin.Admin, error)
	GetAdmin(ctx context.Context, id string) (modeladmin.Admin, error)
	GetAdminAccounts(ctx context.Context) ([]modeladmin.Admin, error)
	LoginAdminAccount(ctx context.Context, email string, password string) (modeladmin.Admin, error)
	DeleteAdminAccount(ctx context.Context, id string) (string, error)
}
