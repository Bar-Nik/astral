package app

import (
	"context"

	"github.com/gofrs/uuid"
)

type (
	// Repo interface for user data repository.
	Repo interface {
		DocsInfoRepo

		UserSave(context.Context, User) (uuid.UUID, error)
		// Delete(context.Context, uuid.UUID) error
		// ByID(context.Context, uuid.UUID) (*User, error)
		// ByEmail(context.Context, string) (*User, error)
		// ByUsername(context.Context, string) (*User, error)
	}
	DocsInfoRepo interface {
		// SaveFile(ctx context.Context) error
		// DeleteFile(ctx context.Context, userID, fileID uuid.UUID) error
		// GetFile(ctx context.Context, fileID uuid.UUID) (*Docs, error)
		// ListFileByUserID(ctx context.Context, userID uuid.UUID) ([]Docs, error)
	}
	// PasswordHash module responsible for hashing password.
	PasswordHash interface {
		// Hashing returns the hashed version of the password.
		// Errors: unknown.
		Hashing(password string) ([]byte, error)
		// Compare compares two passwords for matches.
		Compare(hashedPassword []byte, password []byte) bool
	}

	// FileStore interface for saving and getting files.
	Docs interface {
		// // UploadFile save new file in database.
		// // Errors: unknown.
		// UploadDocs(ctx context.Context, f Docs) (uuid.UUID, error)
		// // DownloadFile get file by id.
		// // Errors: unknown.
		// DownloadDocs(ctx context.Context, id uuid.UUID) (*Docs, error)
		// // DeleteFile delete file by id.
		// // Errors: unknown.
		// DeleteDocs(ctx context.Context, id uuid.UUID) error
	}
	// Sessions module for manager user's session.
	Sessions interface {
		// SessionSave saves the new user session in a database.
		// Errors: unknown.
		SessionSave(context.Context, Session) error
		// SessionByID returns user session by session id.
		// Errors: ErrNotFound, unknown.
		SessionByID(context.Context, uuid.UUID) (*Session, error)
		// SessionDelete removes user session.
		// Errors: ErrNotFound, unknown.
		SessionDelete(context.Context, uuid.UUID) error
	}

	// Auth interface for generate access and refresh token by subject.
	Auth interface {
		// Token generate tokens by subject with expire time.
		// Errors: unknown.
		Token(uuid.UUID) (*Token, error)
		// Subject unwrap Subject info from token.
		// Errors: ErrInvalidToken, ErrExpiredToken, unknown.
		Subject(token string) (uuid.UUID, error)
	}

	// ID generator for session.
	ID interface {
		// New generate new ID for session.
		New() uuid.UUID
	}
)
