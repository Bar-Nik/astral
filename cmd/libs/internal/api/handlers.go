package api

import (
	"backentrymiddle/cmd/libs/internal/app"
	"backentrymiddle/internal/logger"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func (a *api) Register(w http.ResponseWriter, r *http.Request) {
	// userSession := session.FromContext(r.Context())
	// if userSession == nil {
	// 	errorHandler(w, r, http.StatusUnauthorized, Err1)

	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")
	jsong, err := io.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, r, http.StatusUnauthorized, Err2)
		return
	}
	var newUser app.User
	err = json.Unmarshal(jsong, &newUser)
	if err != nil {
		errorHandler(w, r, http.StatusUnauthorized, Err3)
		return
	}

	id, err := a.app.CreateUser(r.Context(), newUser.Email, newUser.Name, string(newUser.PassHash))
	if err != nil {
		errorHandler(w, r, http.StatusUnauthorized, Err4)
		return
	}
	data, err := json.Marshal(id)
	if err != nil {
		errorHandler(w, r, http.StatusUnauthorized, ErrUserUnauthorized)
		return
	}
	w.Write(data)
}

func errorHandler(w http.ResponseWriter, r *http.Request, code int, err error) {
	w.WriteHeader(code)
	erR := json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	if erR != nil {
		logger.FromContext(r.Context()).Error("couldn't send error msg", slog.String(logger.Error.String(), err.Error()))

		return
	}
}
