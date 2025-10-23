package admin

import (
	"context"

	model "github.com/yohannesgossaye/internal/domain/model/admin"
	"github.com/yohannesgossaye/internal/persistence"
)

type AdminService struct {
	adminRepository persistence.AdminRepositary
}

func NewAdminService(adminRepository persistence.AdminRepositary) *AdminService {
	return &AdminService{
		adminRepository: adminRepository,
	}
}

func (s AdminService) CreateAdmin(ctx context.Context, req *model.Admin) (model.Admin, error) {
	return s.adminRepository.CreateAdmin(ctx, req)
}

func (s AdminService) GetAdmin(ctx context.Context, id string) (model.Admin, error) {
	return s.adminRepository.GetAdmin(ctx, id)
}

func (s AdminService) GetAdminAccounts(ctx context.Context) ([]model.Admin, error) {
	return s.adminRepository.GetAdminAccounts(ctx)
}
func (s AdminService) LoginAdminAccount(ctx context.Context, email string, password string) (model.Admin, error) {
	return s.adminRepository.LoginAdminAccount(ctx, email, password)
}

func (s AdminService) DeleteAdminAccount(ctx context.Context, id string) (string, error) {
	return s.adminRepository.DeleteAdminAccount(ctx, id)
}
