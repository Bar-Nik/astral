package repo

import (
	"backentrymiddle/internal/app"
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
	"github.com/golang-migrate/migrate"
)

var _ app.Repo = &Repo{}

type (
	Config struct {
		DSN string `yaml:"dsn"`
	}
	Repo struct {
		db *sql.DB
	}
)

func New(cfg Config) (*Repo, error) {
	migrator, err := migrate.New(
		"file://migrations",
		cfg.DSN)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &Repo{db: conn}, nil
}
func (r *Repo) Close() error {
	return r.db.Close()
}

// Save for implements app.Repo.
func (r *Repo) Save(ctx context.Context, u app.User) (id uuid.UUID, err error) {

	return id, nil
}

// Delete for implements app.Repo.
func (r *Repo) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

// ByID for implements app.Repo.
func (r *Repo) ByID(ctx context.Context, id uuid.UUID) (u *app.User, err error) {

	return u, nil
}

// ByEmail for implements app.Repo.
func (r *Repo) ByEmail(ctx context.Context, email string) (u *app.User, err error) {

	return u, nil
}

// ByUsername for implements app.Repo.
func (r *Repo) ByUsername(ctx context.Context, username string) (u *app.User, err error) {

	return u, nil
}

// SaveAvatar for implements app.Repo.
func (r *Repo) SaveFile(ctx context.Context, userFile app.AvatarInfo) (err error) {
	return nil
}

// DeleteAvatar for implements app.Repo.
func (r *Repo) DeleteFile(ctx context.Context, userID, avatarID uuid.UUID) error {
	return nil
}

// GetAvatar for implements app.Repo.
func (r *Repo) GetFile(ctx context.Context, avatarID uuid.UUID) (f *app.AvatarInfo, err error) {

	return f, nil
}

// ListAvatarByUserID for implements app.Repo.
func (r *Repo) ListFileByUserID(ctx context.Context, userID uuid.UUID) (userAvatars []app.AvatarInfo, err error) {

	return userAvatars, nil
}
