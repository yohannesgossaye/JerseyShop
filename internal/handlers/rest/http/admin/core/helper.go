package core

import (
	"github.com/yohannesgossaye/internal/domain/dto/admin"
	model "github.com/yohannesgossaye/internal/domain/model/admin"
)

func MapcreateAdminRequestToModel(req admin.CreateAdminAccount) model.Admin {
	return model.Admin{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}
