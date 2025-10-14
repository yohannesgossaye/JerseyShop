-- name: CreateAdmin :one
INSERT INTO admins (
    customer_id,
    role,
    permissions,
    current_merchant_numbers
) VALUES (
    $1, $2, $3, $4
)
RETURNING 
    id,
    customer_id,
    role,
    permissions,
    current_merchant_numbers
;

-- name: GetAdmin :one
SELECT * FROM admins WHERE customer_id = $1;

-- name: DeleteAdmin :exec
DELETE FROM admins WHERE customer_id = $1;

-- name: GetAllAdmins :many
SELECT * FROM admins ORDER BY created_at DESC;