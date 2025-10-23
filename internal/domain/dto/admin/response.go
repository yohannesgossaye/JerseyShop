package admin


type AdminResponse struct {
	ID                     int64  `json:"id"`
	CustomerID             int64  `json:"customer_id"`
	Email                  string `json:"email"`
	Password               string `json:"password"`
	Role                   string `json:"role"`
	Permissions            string `json:"permissions"`
	CurrentMerchantNumbers string `json:"current_merchant_numbers"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}

type AdminLoginResponse struct {
	AccessToken string `json:"access_token"`
	Admin       AdminResponse
}

