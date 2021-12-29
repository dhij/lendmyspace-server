package domain

import "time"

type Room struct {
	ID            int64     `json:"id"`
	Name          *string   `json:"name" db:"name"`
	Description   *string   `json:"description" db:"description"`
	HostID        int64     `json:"host_id" db:"host_id"`
	ParticipantID *string   `json:"participant_id" db:"participant_id"`
	MessageID     *string   `json:"message_id" db:"message_id"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
}
