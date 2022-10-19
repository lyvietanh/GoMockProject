-- name: GetUserByEmail :one
SELECT * FROM users
WHERE 
    email=$1
;

-- name: CreateUser :one
INSERT INTO users(
	email, 
    role_id, 
    title, 
    last_name, 
    first_name, 
    date_of_birth, 
    phone, 
    travel_doc_type, 
    travel_doc_number, 
    ffp_number, 
    password_salt, 
    secure_password, 
    enabled, 
    create_date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
) RETURNING *;

-- name: UpdateUserProfile :one
UPDATE users
SET 
    role_id=$2,
    title=$3,
    last_name=$4,
    first_name=$5,
    date_of_birth=$6,
    phone=$7,
    travel_doc_type=$8,
    travel_doc_number=$9,
    ffp_number=$10,
    enabled=$11,
    modify_date=$12
WHERE 
    email=$1
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users
SET 
    password_salt=$2,
    secure_password=$3,
    modify_date=$4
WHERE 
    email=$1
RETURNING *;