-- name: UserSignup :one
INSERT INTO USERS ("email", "password", "nickname") VALUES ($1, $2, $3)
RETURNING id, firstname, lastname, email, password, deleted, nickname, created_at;

-- name: UserGetByEmail :one
SELECT id, firstname, lastname, email, password, deleted, nickname, created_at FROM USERS WHERE email = $1;

-- name: UserGetAll :many
SELECT id, firstname, lastname, email, password, deleted, nickname, created_at FROM USERS;

-- name: UserGetEmails :many
SELECT id, email FROM USERS;

-- name: UserGetById :one
SELECT id, firstname, lastname, email, password, deleted, nickname, created_at FROM USERS WHERE id = $1;

