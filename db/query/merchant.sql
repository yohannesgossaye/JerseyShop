-- name: CreateMerchant :one
INSERT INTO merchants (
  customer_id, store_name, business_license, business_address, business_type,
  business_email, business_phone, owner_first_name, owner_last_name
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetMerchantByCustomerID :one
SELECT * FROM merchants WHERE customer_id = $1;

-- name: VerifyMerchant :exec
UPDATE merchants SET is_verified = TRUE WHERE id = $1;

-- name: GetAllMerchants :many
SELECT * FROM merchants ORDER BY created_at DESC;

-- name: GetAllVerifiedMerchants :many
SELECT * FROM merchants WHERE is_verified = TRUE ORDER BY created_at DESC;

