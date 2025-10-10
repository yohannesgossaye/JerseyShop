package errormessage

import (
	"errors"
	"net/http"
)

var (
	// Customer creation & OTP errors
	ErrCreateCustomer = errors.New("failed to create customer")
	ErrSendOtp        = errors.New("failed to send OTP email")
	ErrGenerateOtp    = errors.New("failed to generate OTP code")
	ErrVerifyCustomer = errors.New("failed to verify customer OTP")
	ErrOtpExpired     = errors.New("OTP has expired")
	ErrInvalidOtp     = errors.New("invalid OTP provided")

	// Login and JWT errors
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrGenerateJwt        = errors.New("failed to generate JWT token")

	// Generic or unexpected errors
	ErrInternalServer = errors.New("internal server error")
	ErrUnexpected     = errors.New("unexpected error")
)

var ErrorMap = map[error]int{
	ErrCreateCustomer:     http.StatusInternalServerError,
	ErrSendOtp:            http.StatusInternalServerError,
	ErrGenerateOtp:        http.StatusInternalServerError,
	ErrVerifyCustomer:     http.StatusInternalServerError,
	ErrOtpExpired:         http.StatusUnauthorized,
	ErrInvalidOtp:         http.StatusUnauthorized,
	ErrInvalidCredentials: http.StatusUnauthorized,
	ErrGenerateJwt:        http.StatusInternalServerError,
	ErrInternalServer:     http.StatusInternalServerError,
	ErrUnexpected:         http.StatusInternalServerError,
}
