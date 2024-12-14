package repo

import (
	"backentrymiddle/cmd/libs/internal/app"
	"time"

	"github.com/gofrs/uuid"
)

type (
	user struct {
		ID       uuid.UUID `db:"id" json:"id"`
		Email    string    `db:"email" json:"email"`
		Name     string    `db:"name" json:"name"`
		PassHash []byte    `db:"pass_hash" json:"pass_hash"`
		//Status    string    `db:"status" json:"status"`
		CreatedAt time.Time `db:"created_at" json:"created_at"`
		UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	}
)

func convert(u app.User) *user {
	return &user{
		ID:       u.ID,
		Email:    u.Email,
		Name:     u.Name,
		PassHash: u.PassHash,
		//Status:    u.Status.String(),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// func (u user) convert() *app.User {
// 	return &app.User{
// 		ID:       u.ID,
// 		Email:    u.Email,
// 		Name:     u.Name,
// 		PassHash: u.PassHash,
// 		//Status:    appUserStatus(u.Status),
// 		CreatedAt: u.CreatedAt,
// 		UpdatedAt: u.UpdatedAt,
// 	}
// }
