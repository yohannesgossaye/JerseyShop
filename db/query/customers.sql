-- name: CreateCustomer :one
INSERT INTO customers (
    first_name,
    last_name,
    email,
    phone_number,
    password_hash,
    country,
    city,
    address,
    zip_code,
    gender,
    date_of_birth,
    otp_code,
    otp_expires_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING 
    id,
    first_name,
    last_name,
    email,
    phone_number,
    country,
    city,
    address,
    zip_code,
    gender,
    date_of_birth,
    otp_code,
    otp_expires_at,
    is_active,
    is_admin,
    created_at,
    updated_at;

-- name: VerifyCustomerOTP :one
UPDATE customers
SET 
    is_active = TRUE,
    otp_code = NULL,
    otp_expires_at = NULL,
    updated_at = NOW()
WHERE 
    email = $1
    AND otp_code = $2
    AND otp_expires_at > NOW()
RETURNING 
    id, email, is_active, updated_at;
    
-- name: GetCustomerByEmail :one
SELECT 
    id,
    first_name,
    last_name,
    email,
    password_hash,
    is_active,
    is_admin
FROM customers
WHERE email = $1
LIMIT 1;
-- name: UpdateCustomerOTP :exec
UPDATE customers
SET 
    otp_code = $1,
    otp_expires_at = $2,
    updated_at = NOW()
WHERE email = $3;
