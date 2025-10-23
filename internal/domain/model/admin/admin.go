package admin

type Admin struct {
	ID                     int64  `json:"id"`
	CustomerID             int64  `json:"customer_id"`
	Role                   string `json:"role"`
	Password               string `json:"password"`
	Email                  string `json:"email"`
	Permissions            string `json:"permissions"`
	CurrentMerchantNumbers string `json:"current_merchant_numbers"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}
