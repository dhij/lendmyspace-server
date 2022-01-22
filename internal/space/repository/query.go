package repository

const (
	GetSpaceQuery    = `SELECT * FROM spaces WHERE id = $1 LIMIT 1`
	ListSpacesQuery  = `SELECT id, name, description, location, link FROM spaces ORDER BY id`
	CreateSpaceQuery = `INSERT INTO spaces (name, description, location, link, host_id, images, dates) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
)
