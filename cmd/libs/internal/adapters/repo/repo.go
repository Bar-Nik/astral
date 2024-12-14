package repo

import (
	"backentrymiddle/cmd/libs/internal/app"
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sipki-tech/database"
	"github.com/sipki-tech/database/connectors"
	"github.com/sipki-tech/database/migrations"
)

var _ app.Repo = &Repo{}

type (
	Config struct {
		DSN        connectors.Raw
		MigrateDir string
		Driver     string
	}
	Repo struct {
		db *database.SQL
	}
)

func New(ctx context.Context, cfg Config) (*Repo, error) {
	migrates, err := migrations.Parse(cfg.MigrateDir)
	if err != nil {
		return nil, fmt.Errorf("migrations.Parse: %w", err)
	}
	err = migrations.Run(ctx, cfg.Driver, &cfg.DSN, migrations.Up, migrates)
	if err != nil {
		return nil, fmt.Errorf("migrations.Run: %w", err)
	}

	conn, err := database.NewSQL(ctx, cfg.Driver, database.SQLConfig{}, &cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}
	return &Repo{db: conn}, nil
}
func (r *Repo) Close() error {
	return r.db.Close()
}

// Save for implements app.Repo.
func (r *Repo) UserSave(ctx context.Context, u app.User) (id uuid.UUID, err error) {
	err = r.db.NoTx(func(db *sqlx.DB) error {
		newUser := convert(u)
		const query = "INSERT INTO users (email, name pass_hash) VALUES($1, $2, $3) RETURNING id"

		err := db.GetContext(ctx, &id, query, newUser.Email, newUser.Name, newUser.PassHash)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", err)
		}

		return nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
	// query := "INSERT INTO users (email, name pass_hash) VALUES($1, $2, $3) RETURNING id"
	// newUser := convert(u)
	// err = r.db.QueryRowContext(ctx, query, newUser.Email, newUser.Name, newUser.PassHash).Scan(&newUser.ID, &newUser.Email, &newUser.Name, &newUser.PassHash, &newUser.CreatedAt, &newUser.UpdatedAt)
	// if err != nil {
	// 	return uuid.Nil, err
	// }

	// return newUser.ID, nil
}

// func (d Repository) SaveBookToDatabase(book domain.Book, ctx context.Context) (domain.Book, error) {
// 	query := "INSERT INTO books (title, year_book) VALUES($1,$2) RETURNING *"
// 	err := d.db.QueryRowContext(ctx, query, book.Title, book.Year).Scan(&book.ID, &book.Title, &book.Year)
// 	if err != nil {
// 		return domain.Book{}, err
// 	}
// 	return book, nil

// }

// Delete for implements app.Repo.
// func (r *Repo) Delete(ctx context.Context, id uuid.UUID) error {
// 	return nil
// }

// // ByID for implements app.Repo.
// func (r *Repo) ByID(ctx context.Context, id uuid.UUID) (u *app.User, err error) {

// 	return u, nil
// }

// // ByEmail for implements app.Repo.
// func (r *Repo) ByEmail(ctx context.Context, email string) (u *app.User, err error) {

// 	return u, nil
// }

// // ByUsername for implements app.Repo.
// func (r *Repo) ByUsername(ctx context.Context, username string) (u *app.User, err error) {

// 	return u, nil
// }

// // SaveAvatar for implements app.Repo.
// func (r *Repo) SaveFile(ctx context.Context, userFile app.Docs) (err error) {
// 	return nil
// }

// // DeleteAvatar for implements app.Repo.
// func (r *Repo) DeleteFile(ctx context.Context, userID, avatarID uuid.UUID) error {
// 	return nil
// }

// // GetAvatar for implements app.Repo.
// func (r *Repo) GetFile(ctx context.Context, avatarID uuid.UUID) (f *app.Docs, err error) {

// 	return f, nil
// }

// // ListAvatarByUserID for implements app.Repo.
// func (r *Repo) ListFileByUserID(ctx context.Context, userID uuid.UUID) (userAvatars []app.Docs, err error) {

// 	return userAvatars, nil
// }
