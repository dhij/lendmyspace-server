package repository

const (
	GetSpaceQuery    = `SELECT * FROM spaces WHERE id = $1 LIMIT 1`
	CreateSpaceQuery = `INSERT INTO spacess (name, description, image, host_id, date_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
)
