package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	dtoadmin "github.com/yohannesgossaye/internal/domain/dto/admin"
	core "github.com/yohannesgossaye/internal/handlers/rest/http/admin/core"
	"github.com/yohannesgossaye/internal/service"
	response "github.com/yohannesgossaye/pkgs/message/Response"
	errormessage "github.com/yohannesgossaye/pkgs/message/errormessage"
	"github.com/yohannesgossaye/pkgs/message/localization"

	logger "github.com/yohannesgossaye/pkgs/logger"
)

type AdminHandler struct {
	adminService service.AdminAccountService
	logger       logger.Logger
}

func NewAdminHandler(adminService service.AdminAccountService, logger logger.Logger) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
		logger:       logger,
	}
}

func (h *AdminHandler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("received request to create admin")
	var req dtoadmin.CreateAdminAccount
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Errorf("Error decoding request body: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrAdminCreateDecodeFailed.Code)
		return
	}
	if err := req.Validate(); err != nil {
		h.logger.Errorf("Error validating request: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrorValidationFailed.Code)
		return
	}
	CreateMap := core.MapcreateAdminRequestToModel(req)
	h.logger.Infof("🔍 Calling service with mapped data: %+v", CreateMap)
	admin, err := h.adminService.CreateAdmin(r.Context(), &CreateMap)
	if err != nil {
		h.logger.Errorf("Error creating admin: %v", err)
		if errors.Is(err, errormessage.ErrEmailDuplicate) {
			localization.SendErrorByCodeResponse(w, localization.ErrAdminEmailAlreadyExists.Code)
			return
		}
		switch err.Error() {
		case localization.ErrAdminEmailAlreadyExists.Message:
			localization.SendErrorByCodeResponse(w, localization.ErrAdminEmailAlreadyExists.Code)
		default:
			localization.SendErrorByCodeResponse(w, localization.ErrAdminCreate.Code)
		}
		return
	}
	response.SendSuccessResponse(w, http.StatusOK, "Admin created successfully", admin, nil)
}

func (h *AdminHandler) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("receiving login request ")
	var req dtoadmin.LoginAdminAccount
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Errorf("Error decoding request body: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrAdminCreateDecodeFailed.Code)
		return
	}
	if err := req.Validate(); err != nil {
		h.logger.Errorf("Error validating request: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrorValidationFailed.Code)
		return
	}
	admin, err := h.adminService.LoginAdminAccount(r.Context(), req.Email, req.Password)
	if err != nil {
		h.logger.Errorf("Error logging in admin: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrAdminLogin.Code)
		return
	}
	response.SendSuccessResponse(w, http.StatusOK, "Admin logged in successfully", admin, nil)
}

func (h *AdminHandler) GetAdmin(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("receiving get admin request ")
	id := r.URL.Query().Get("id")
	admin, err := h.adminService.GetAdmin(r.Context(), id)
	if err != nil {
		h.logger.Errorf("Error getting admin: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrAdminGet.Code)
		return
	}
	response.SendSuccessResponse(w, http.StatusOK, "Admin retrieved successfully", admin, nil)
}

func (h *AdminHandler) GetAdminAccounts(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("receiving get admin accounts request ")
	admins, err := h.adminService.GetAdminAccounts(r.Context())
	if err != nil {
		h.logger.Errorf("Error getting admin accounts: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrAdminGet.Code)
		return
	}
	response.SendSuccessResponse(w, http.StatusOK, "Admin accounts retrieved successfully", admins, nil)
}

func (h *AdminHandler) DeleteAdminAccount(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("receiving delete admin account ")
	id := r.URL.Query().Get("id")
	_, err := h.adminService.DeleteAdminAccount(r.Context(), id)
	if err != nil {
		h.logger.Errorf("Error deleting admin account: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrAdminDelete.Code)
		return
	}
	response.SendSuccessResponse(w, http.StatusOK, "Admin account deleted successfully", nil, nil)
}
