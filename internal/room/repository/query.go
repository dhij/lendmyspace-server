package repository

const (
	GetRoomQuery    = `SELECT * FROM rooms WHERE id = $1 LIMIT 1`
	CreateRoomQuery = `INSERT INTO rooms (name, description, host_id, participant_id, message_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
)
