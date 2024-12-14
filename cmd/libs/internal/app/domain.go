package app

import (
	"io"
	"net"
	"time"

	"github.com/gofrs/uuid"
)

type (
	User struct {
		ID       uuid.UUID
		Email    string
		Name     string
		PassHash []byte
		//Status    dom.UserStatus
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	File struct {
		ID          uuid.UUID
		UserID      uuid.UUID
		Name        string
		ContentType string
		Size        int64
		ModTime     time.Time
		io.ReadSeekCloser
	}

	Token struct {
		// Generate by Auth contract.
		Value string
	}

	Origin struct {
		IP        net.IP
		UserAgent string
	}

	Session struct {
		ID        uuid.UUID // Generate by repository layer.
		Origin    Origin
		Token     Token
		UserID    uuid.UUID
		CreatedAt time.Time // Generate by repository layer.
		UpdatedAt time.Time // Generate by repository layer.
	}
)
