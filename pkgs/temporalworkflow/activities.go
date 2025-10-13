package temporalworkflow

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	podb "github.com/yohannesgossaye/internal/persistence/postgres"
	"github.com/yohannesgossaye/pkgs/utils/email"
)

// Activity 1: Save user to DB
func SaveUserActivity(ctx context.Context, input UserWorkflowInput) (string, error) {

	connStr := getenv("DATABASE_URL", "")
	if connStr == "" {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	}
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return "", err
	}
	defer pool.Close()

	repo := podb.NewCustomerRepositary(pool)
	user := MapWorkflowToModel(input)

	// ensure OTP expiry
	if user.OtpCode != "" && user.OtpExpiresAt.IsZero() {
		user.OtpExpiresAt = time.Now().Add(time.Hour)
	}

	fmt.Printf("🔍 Mapped user model: %+v\n", user)
	created, err := repo.CreateCustomer(ctx, &user)
	if err != nil {

		return "", err
	}

	fmt.Printf("✅ Customer created successfully with ID: %d\n", created.ID)
	return fmt.Sprintf("%d", created.ID), nil
}

// Activity 2: Send email with OTP
func SendEmailActivity(ctx context.Context, emailAddr string, name string, otp string) error {
	subject := "🎉 Welcome to Our Service"
	body := fmt.Sprintf(
		"Hi %s,\n\nThanks for registering with our service!\n\nYour OTP is: %s\nThis OTP will expire in 1 hour.\n\nBest regards,\nJosambin Sports Jersey Shop",
		name, otp,
	)

	sender := email.NewSMTPSender()
	if sender == nil {
		return fmt.Errorf(" failed to initialize SMTP sender — missing env vars")
	}

	fmt.Printf("📧 Sending email to %s\n%s\n", emailAddr, body)
	if err := sender.Send(emailAddr, subject, body); err != nil {
		return err
	}

	fmt.Println("✅ Email sent successfully!")
	return nil
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
