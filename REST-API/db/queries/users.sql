



-- name: CreateUser :one

INSERT INTO users(
    first_name,
    last_name,
    email,
    phone_number
)VALUES(
    $1,
    $2,
    $3,
    $4
)RETURNING id;

-- name: ListUsers :many
SELECT * FROM users
OFFSET $1 LIMIT $2;


-- name: GetUser :one
SELECT * FROM users WHERE id = $1;


-- name: UpdateUser :exec
UPDATE users SET
    first_name = $2,
    last_name = $3,
    email = $4,
    phone_number = $5
WHERE id = $1;


-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1 ;
