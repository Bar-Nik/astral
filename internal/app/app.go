package app

// App manages business logic methods.
type App struct {
	repo Repo
	docs Docs
}

// New build and returns new App.
func New(r Repo, d Docs) *App {
	return &App{
		repo: r,
		docs: d,
	}
}
