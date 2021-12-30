package repository

const (
	GetUserQuery    = `SELECT user_name FROM users WHERE id = $1 LIMIT 1`
	ListUsersQuery  = `SELECT user_name, first_name, last_name, email FROM users ORDER BY id LIMIT $1 OFFSET $2`
	CreateUserQuery = `INSERT INTO users(user_name, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id`
)
