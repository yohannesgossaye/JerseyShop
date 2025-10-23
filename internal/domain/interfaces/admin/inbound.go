package admin

import "net/http"

type AdminAccountHandler interface {
	CreateAdminAccount(w http.ResponseWriter, r *http.Request)
	LoginAdminAccount(w http.ResponseWriter, r *http.Request)
	GetAdminAccount(w http.ResponseWriter, r *http.Request)
	GetAdminAccounts(w http.ResponseWriter, r *http.Request)
	DeleteAdminAccount(w http.ResponseWriter, r *http.Request)
}
