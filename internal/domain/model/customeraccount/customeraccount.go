package customeraccount

import "time"

type Customer struct {
	ID           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Password     string    `json:"-"`
	Country      string    `json:"country"`
	City         string    `json:"city"`
	Address      string    `json:"address"`
	ZipCode      string    `json:"zip_code"`
	Gender       string    `json:"gender"`
	DateOfBirth  string    `json:"date_of_birth"`
	IsActive     bool      `json:"is_active"`
	IsAdmin      bool      `json:"is_admin"`
	OtpCode      string    `json:"-"`
	OtpExpiresAt time.Time `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
