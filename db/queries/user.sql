-- name:CreateUser:one
INSERT INTO user(
    user_id,
    created_at,
    updated_at,
    name,
    phoneNumber,
    email,
    address,
    password
)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8);
RETURNING *;