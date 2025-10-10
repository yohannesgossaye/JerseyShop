package customeraccount

import (
	"context"
	"errors"
	"time"

	"fmt"
	"strings"

	custmeraccountdto "github.com/yohannesgossaye/internal/domain/dto/customeraccount"
	model "github.com/yohannesgossaye/internal/domain/model/customeraccount"
	"github.com/yohannesgossaye/internal/persistence"
	errormessage "github.com/yohannesgossaye/pkgs/message/errormessage"
	"github.com/yohannesgossaye/pkgs/message/localization"
	"github.com/yohannesgossaye/pkgs/utils/email"
	helper "github.com/yohannesgossaye/pkgs/utils/helper"
)

type CustomerService struct {
	repo        persistence.CustomerRepositary
	emailSender email.Sender
}

func NewCustomerService(repo persistence.CustomerRepositary, emailSender email.Sender) *CustomerService {
	return &CustomerService{
		repo:        repo,
		emailSender: emailSender,
	}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req model.Customer) (model.Customer, error) {
	// generate OTP
	otp := helper.GenerateOtp()
	req.OtpCode = otp
	req.OtpExpiresAt = time.Now().Add(time.Hour)

	// create customer in DB
	createdCustomer, err := s.repo.CreateCustomer(ctx, &req)
	if err != nil {

		if strings.Contains(strings.ToLower(err.Error()), "duplicate") || strings.Contains(err.Error(), "23505") {
			return model.Customer{}, errormessage.ErrEmailDuplicate
		}
		return model.Customer{}, err
	}

	// send welcome email and OTP asynchronously
	go func(emailTo string, fullName string, otp string) {
		body := fmt.Sprintf(
			"Hi %s,\n\nThanks for registering with our service!\n\nYour OTP is: %s\nThis OTP will expire in 1 hour.",
			fullName, otp,
		)
		if err := s.emailSender.Send(emailTo, "Welcome 🎉", body); err != nil {
			fmt.Printf("❌ Failed to send email to %s: %v\n", emailTo, err)
		} else {
			fmt.Printf("📧 Welcome email sent to %s\n", emailTo)
		}
	}(createdCustomer.Email, createdCustomer.FirstName+" "+createdCustomer.LastName, otp)

	return createdCustomer, nil
}

func (s *CustomerService) LoginCustomer(ctx context.Context, email string, password string) (model.Customer, error) {

	customer, err := s.repo.LoginCustomer(ctx, email, password)
	if err != nil {
		return model.Customer{}, errors.New(localization.ErrInvalidCredentials.Message)
	}
	// ensure account is active
	if !customer.IsActive {
		return model.Customer{}, errors.New(localization.ErrCustomerNotActive.Message)
	}
	return customer, nil
}

func (s *CustomerService) VerifyCustomerOtp(ctx context.Context, email, otpCode string) (custmeraccountdto.VerifyOtpResponse, error) {
	customer, err := s.repo.VerifyCustomerOtp(ctx, email, otpCode)
	if err != nil {

		switch err {
		case errormessage.ErrInvalidOtp:
			return custmeraccountdto.VerifyOtpResponse{}, errors.New(localization.ErrInvalidOtp.Message)
		case errormessage.ErrNotExistOtp:
			return custmeraccountdto.VerifyOtpResponse{}, errors.New(localization.ErrNotCorrectOtp.Message)
		case errormessage.ErrOtpExpired:
			return custmeraccountdto.VerifyOtpResponse{}, errors.New(localization.ErrOtpExpired.Message)
		case errormessage.ErrAlreadyVerified:
			return custmeraccountdto.VerifyOtpResponse{}, errors.New(localization.ErrCustomerAlreadyVerified.Message)
		default:

			return custmeraccountdto.VerifyOtpResponse{}, err
		}
	}
	// Just issue JWT and return the response.
	token, err := helper.GenerateJWT(customer.ID, customer.Email)
	if err != nil {
		return custmeraccountdto.VerifyOtpResponse{}, err
	}

	return custmeraccountdto.VerifyOtpResponse{
		AccessToken: token,
		Customer:    customer,
		// Message:     "Account verified successfully",
	}, nil
}
