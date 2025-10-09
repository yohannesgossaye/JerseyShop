package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"github.com/yohannesgossaye/internal/domain/model/customeraccount"
	"github.com/yohannesgossaye/internal/persistence/postgres/gen"
)

type CustomerRepositary struct {
	db  *pgxpool.Pool
	gen *gen.Queries
}

func NewCustomerRepositary(db *pgxpool.Pool) *CustomerRepositary {
	return &CustomerRepositary{
		db:  db,
		gen: gen.New(db),
	}
}

func (r *CustomerRepositary) CreateCustomer(ctx context.Context, req *customeraccount.Customer) (customeraccount.Customer, error) {
	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return customeraccount.Customer{}, err
	}

	// parse date of birth (expecting YYYY-MM-DD)
	var dob pgtype.Date
	if req.DateOfBirth != "" {
		t, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err != nil {
			return customeraccount.Customer{}, err
		}
		dob = pgtype.Date{Time: t, Valid: true}
	}

	// prepare otp pointer
	var otpPtr *string
	if req.OtpCode != "" {
		otpPtr = &req.OtpCode
	}

	customer, err := r.gen.CreateCustomer(ctx, &gen.CreateCustomerParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		PasswordHash: string(hashed),
		Country:      req.Country,
		City:         req.City,
		Address:      req.Address,
		ZipCode:      req.ZipCode,
		Gender:       req.Gender,
		DateOfBirth:  dob,
		OtpCode:      otpPtr,
		OtpExpiresAt: req.OtpExpiresAt,
	})
	if err != nil {
		return customeraccount.Customer{}, err
	}
	// map returned row to domain model
	return customeraccount.Customer{
		ID:          int64(customer.ID),
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		IsActive:    customer.IsActive,
		IsAdmin:     customer.IsAdmin,
		CreatedAt:   customer.CreatedAt,
		UpdatedAt:   customer.UpdatedAt,
	}, nil
}

func (r *CustomerRepositary) LoginCustomer(ctx context.Context, email string, password string) (customeraccount.Customer, error) {
	row, err := r.gen.GetCustomerByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customeraccount.Customer{}, fmtError("invalid credentials")
		}
		return customeraccount.Customer{}, err
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(row.PasswordHash), []byte(password)); err != nil {
		return customeraccount.Customer{}, fmtError("invalid credentials")
	}

	return customeraccount.Customer{
		ID:        int64(row.ID),
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Email:     row.Email,
		IsActive:  row.IsActive,
		IsAdmin:   row.IsAdmin,
	}, nil
}

func (r *CustomerRepositary) VerifyCustomerOtp(ctx context.Context, email string, otpCode string) (customeraccount.Customer, error) {
	// fetch full customer record to check otp and expiry
	var (
		id                                                                                               int
		firstName, lastName, emailDB, phoneNumber, passwordHash, country, city, address, zipCode, gender sql.NullString
		dateOfBirth                                                                                      pgtype.Date
		otpCodePtr                                                                                       sql.NullString
		otpExpiresAt                                                                                     sql.NullTime
		isActive, isAdmin                                                                                bool
		createdAt, updatedAt                                                                             time.Time
	)

	query := `SELECT id, first_name, last_name, email, phone_number, password_hash, country, city, address, zip_code, gender, date_of_birth, otp_code, otp_expires_at, is_active, is_admin, created_at, updated_at FROM customers WHERE email = $1 LIMIT 1`
	row := r.db.QueryRow(ctx, query, email)
	if err := row.Scan(
		&id,
		&firstName,
		&lastName,
		&emailDB,
		&phoneNumber,
		&passwordHash,
		&country,
		&city,
		&address,
		&zipCode,
		&gender,
		&dateOfBirth,
		&otpCodePtr,
		&otpExpiresAt,
		&isActive,
		&isAdmin,
		&createdAt,
		&updatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customeraccount.Customer{}, fmtError("invalid otp or email")
		}
		return customeraccount.Customer{}, err
	}

	if !otpCodePtr.Valid || otpCodePtr.String != otpCode {
		return customeraccount.Customer{}, fmtError("invalid otp")
	}
	if !otpExpiresAt.Valid || time.Now().After(otpExpiresAt.Time) {
		return customeraccount.Customer{}, fmtError("otp expired")
	}

	// mark verified using generated query
	if _, err := r.gen.VerifyCustomerOTP(ctx, &gen.VerifyCustomerOTPParams{Email: email, OtpCode: &otpCode}); err != nil {
		return customeraccount.Customer{}, err
	}

	// return mapped domain customer
	cust := customeraccount.Customer{
		ID:          int64(id),
		FirstName:   firstName.String,
		LastName:    lastName.String,
		Email:       emailDB.String,
		PhoneNumber: phoneNumber.String,
		IsActive:    true,
		IsAdmin:     isAdmin,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
	if otpExpiresAt.Valid {
		cust.OtpExpiresAt = otpExpiresAt.Time
	}
	return cust, nil
}
func (r *CustomerRepositary) UpdateCustomer(ctx context.Context, customer *customeraccount.Customer) error {
	_, err := r.db.Exec(ctx, `
		UPDATE customers
		SET 
			is_active = $1,
			otp_code = $2,
			updated_at = NOW()
		WHERE id = $3
	`, customer.IsActive, customer.OtpCode, customer.ID)

	if err != nil {
		return err
	}

	return nil
}

// small helper to return formatted error (avoids importing fmt repeatedly)
func fmtError(msg string) error {
	return errors.New(msg)
}
