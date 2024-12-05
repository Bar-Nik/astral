package app

import (
	"time"

	"github.com/gofrs/uuid"
)

type (
	User struct {
		ID       uuid.UUID
		Email    string
		Name     string
		PassHash []byte
		DocsID   uuid.UUID
		// Status    dom.UserStatus
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
