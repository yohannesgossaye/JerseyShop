package admin

type CreateAdminAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginAdminAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
