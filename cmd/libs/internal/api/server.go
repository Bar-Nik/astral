package api

import (
	"backentrymiddle/cmd/libs/internal/app"
	"backentrymiddle/internal/logger"
	"context"
	"errors"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

var (
	ErrUserUnauthorized           = errors.New("user unauthorized")
	ErrMissingAuthorizationHeader = errors.New("missing authorization header")
	ErrBadAuthorizationString     = errors.New("bad authorization string")
	ErrInvalidArgument            = errors.New("invalid argument")
	ErrMaxAvatarSize              = errors.New("max file size 25 mb")
	Err1                          = errors.New("1")
	Err2                          = errors.New("2")
	Err3                          = errors.New("3")
	Err4                          = errors.New("4")
)

type application interface {
	CreateUser(ctx context.Context, email, username, password string) (uuid.UUID, error)
}

type api struct {
	app application
}

func New(ctx context.Context, applications *app.App) http.Handler {
	log := logger.FromContext(ctx)

	api := api{
		app: applications,
	}

	router := mux.NewRouter()
	router.Use(Logging(log))

	router.HandleFunc("/api/register", api.Register).Methods(http.MethodPost)
	// router.HandleFunc("/api/auth", ourServer.Auth).Methods(http.MethodPost)
	// router.HandleFunc("/api/docs", ourServer.UploadDocs).Methods(http.MethodPost)
	// router.HandleFunc("/api/docs", ourServer.GetAllDocs).Methods(http.MethodGet)
	// router.HandleFunc("/api/docs/{id}", ourServer.GetDocs).Methods(http.MethodGet)
	// router.HandleFunc("/api/docs/{id}", ourServer.DeleteDocs).Methods(http.MethodDelete)
	// router.HandleFunc("/api/auth/", ourServer.CloseSession).Methods(http.MethodDelete)

	// err = http.ListenAndServe("127.0.0.1:8080", r)
	// if err != nil {
	// 	log.Debug("Server failed")
	// }

	return router
}

// type Server struct {
// 	Datebase *repo.Repo
// }

// func (s Server) Register(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func (s Server) Auth(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func (s Server) UploadDocs(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func (s Server) GetAllDocs(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func (s Server) GetDocs(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func (s Server) DeleteDocs(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func (s Server) CloseSession(w http.ResponseWriter, r *http.Request) {
// 	return
// }
