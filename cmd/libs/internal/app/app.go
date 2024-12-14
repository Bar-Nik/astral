package app

// App manages business logic methods.
type App struct {
	repo Repo
	docs Docs
	// sessions Sessions
	hash PasswordHash
}

// New build and returns new App.
func New(r Repo, d Docs, ph PasswordHash) *App {
	return &App{
		repo: r,
		docs: d,
	}
}
