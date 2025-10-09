package customeraccount

type CreateCustomer struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
}

type VerifyCustomerOtp struct {
	Email   string `json:"email"`
	OTPCode string `json:"otp_code"`
}
type LoginCustomer struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
