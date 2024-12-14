package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

func (a *App) CreateUser(ctx context.Context, email, username, password string) (userID uuid.UUID, err error) {
	passHash, err := a.hash.Hashing(password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("a.hash.Hashing: %w", err)
	}
	email = strings.ToLower(email)
	newUser := User{
		Email:    email,
		Name:     username,
		PassHash: passHash,
	}
	userID, err = a.repo.UserSave(ctx, newUser)
	if err != nil {
		return uuid.Nil, fmt.Errorf("repo.UserSave: %w", err)
	}

	return userID, nil
}
