package app

import (
	"context"

	"github.com/gofrs/uuid"
)

type (
	// Repo interface for user data repository.
	Repo interface {
		DocsInfoRepo

		Save(context.Context, User) (uuid.UUID, error)
		Delete(context.Context, uuid.UUID) error
		ByID(context.Context, uuid.UUID) (*User, error)
		ByEmail(context.Context, string) (*User, error)
		ByUsername(context.Context, string) (*User, error)
	}
	DocsInfoRepo interface {
		SaveFile(ctx context.Context, fileCache AvatarInfo) error
		DeleteFile(ctx context.Context, userID, fileID uuid.UUID) error
		GetFile(ctx context.Context, fileID uuid.UUID) (*AvatarInfo, error)
		ListFileByUserID(ctx context.Context, userID uuid.UUID) ([]AvatarInfo, error)
	}

	// FileStore interface for saving and getting files.
	Docs interface {
		// UploadFile save new file in database.
		// Errors: unknown.
		UploadDocs(ctx context.Context, f Avatar) (uuid.UUID, error)
		// DownloadFile get file by id.
		// Errors: unknown.
		DownloadDocs(ctx context.Context, id uuid.UUID) (*Avatar, error)
		// DeleteFile delete file by id.
		// Errors: unknown.
		DeleteDocs(ctx context.Context, id uuid.UUID) error
	}
)
