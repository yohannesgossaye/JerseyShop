package helper

import (
	"crypto/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateOtp generates a secure 6-digit OTP
func GenerateOtp() string {
	const otpLength = 6
	const digits = "0123456789"

	otp := make([]byte, otpLength)
	_, err := rand.Read(otp)
	if err != nil {
		panic(err)
	}

	for i := 0; i < otpLength; i++ {
		otp[i] = digits[int(otp[i])%len(digits)]
	}

	return string(otp)
}

func GenerateJWT(userID int64, email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "defaultsecret"
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
