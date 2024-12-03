package api

import (
	"backentrymiddle/internal/db"
	"net/http"
)

type Server struct {
	Datebase *db.Repository
}

func (s Server) Register(w http.ResponseWriter, r *http.Request) {
	return
}

func (s Server) Auth(w http.ResponseWriter, r *http.Request) {
	return
}

func (s Server) UploadDocs(w http.ResponseWriter, r *http.Request) {
	return
}

func (s Server) GetAllDocs(w http.ResponseWriter, r *http.Request) {
	return
}

func (s Server) GetDocs(w http.ResponseWriter, r *http.Request) {
	return
}

func (s Server) DeleteDocs(w http.ResponseWriter, r *http.Request) {
	return
}

func (s Server) CloseSession(w http.ResponseWriter, r *http.Request) {
	return
}
