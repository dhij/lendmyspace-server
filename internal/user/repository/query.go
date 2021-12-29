package repository

const (
	FindByIDQuery   = `SELECT user_name FROM users WHERE id = $1`
	CreateUserQuery = `INSERT INTO users(user_name, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id`
)
