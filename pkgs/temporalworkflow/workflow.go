package temporalworkflow

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type UserWorkflowInput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	Gender      string `json:"gender"`
	OtpCode     string `json:"otp_code"`
}

func UserSignupWorkflow(ctx workflow.Context, input UserWorkflowInput) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts:    5,
			InitialInterval:    time.Second * 2,
			BackoffCoefficient: 2.0,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var userID string
	if err := workflow.ExecuteActivity(ctx, SaveUserActivity, input).Get(ctx, &userID); err != nil {
		return "", err
	}
	workflow.GetLogger(ctx).Info("🕒 Waiting 15 seconds before sending confirmation email...")
	if err := workflow.Sleep(ctx, 15*time.Second); err != nil {
		return "", err
	}

	if err := workflow.ExecuteActivity(ctx, SendEmailActivity,
		input.Email,
		input.FirstName+" "+input.LastName,
		input.OtpCode,
	).Get(ctx, nil); err != nil {
		return "", err
	}

	return userID, nil
}
