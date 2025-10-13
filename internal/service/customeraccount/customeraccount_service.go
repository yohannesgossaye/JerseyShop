package customeraccount

import (
	"context"
	"errors"
	"strings"
	"time"

	"fmt"
	"os"

	custmeraccountdto "github.com/yohannesgossaye/internal/domain/dto/customeraccount"
	model "github.com/yohannesgossaye/internal/domain/model/customeraccount"
	"github.com/yohannesgossaye/internal/persistence"
	"github.com/yohannesgossaye/internal/service/customeraccount/core"
	errormessage "github.com/yohannesgossaye/pkgs/message/errormessage"
	"github.com/yohannesgossaye/pkgs/message/localization"
	workflows "github.com/yohannesgossaye/pkgs/temporalworkflow"
	"github.com/yohannesgossaye/pkgs/utils/email"
	helper "github.com/yohannesgossaye/pkgs/utils/helper"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
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
	fmt.Printf(" Service received customer: %+v\n", req)

	// Generate OTP for account verification
	otp := helper.GenerateOtp()
	req.OtpCode = otp
	req.OtpExpiresAt = time.Now().Add(time.Hour)

	hostPort := os.Getenv("TEMPORAL_HOSTPORT")
	if hostPort == "" {
		hostPort = "127.0.0.1:7233"
	}
	// Create Temporal client
	c, err := client.NewClient(client.Options{HostPort: hostPort})
	if err != nil {
		return model.Customer{}, err
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:                    "user_signup_" + req.Email,
		TaskQueue:             "USER_TASK_QUEUE_V2",
		WorkflowIDReusePolicy: enumspb.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
	}

	input := core.Toworkflowinput(req)
	fmt.Printf(" Starting workflow asynchronously with input: %+v\n", input)

	// Start workflow **without waiting for completion**
	we, err := c.ExecuteWorkflow(ctx, workflowOptions, workflows.UserSignupWorkflow, input)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") || strings.Contains(err.Error(), "23505") {
			return model.Customer{}, errormessage.ErrEmailDuplicate
		}
		return model.Customer{}, err
	}

	fmt.Println(" Workflow started asynchronously", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Immediately return the customer object to API caller
	return req, nil
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
