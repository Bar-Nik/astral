package main

import (
	"backentrymiddle/internal/api"
	"backentrymiddle/internal/db"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DSN      string `yaml:"dsn"`
	LogLevel int    `yaml:"loglevel"`
}

func main() {
	yamlContent, err := os.ReadFile("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var systemConfig Config
	err = yaml.Unmarshal(yamlContent, &systemConfig)
	if err != nil {
		log.Fatal(err)
	}

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(systemConfig.LogLevel),
	}))

	r := mux.NewRouter()

	migrator, err := migrate.New(
		"file://migrations",
		systemConfig.DSN)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		os.Exit(1)
	}

	rowSQLConn, err := sql.Open("postgres", systemConfig.DSN)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repo := db.NewRepository(rowSQLConn)
	r.Use(api.Logging(log))

	ourServer := api.Server{
		Datebase: repo,
	}
	r.HandleFunc("/api/register", ourServer.Register).Methods(http.MethodPost)
	r.HandleFunc("/api/auth", ourServer.Auth).Methods(http.MethodPost)
	r.HandleFunc("/api/docs", ourServer.UploadDocs).Methods(http.MethodPost)
	r.HandleFunc("/api/docs", ourServer.GetAllDocs).Methods(http.MethodGet)
	// r.HandleFunc("/api/docs", ourServer.GetDocs).Methods(http.MethodGet)
	r.HandleFunc("/api/docs", ourServer.DeleteDocs).Methods(http.MethodDelete)
	r.HandleFunc("/api/auth", ourServer.CloseSession).Methods(http.MethodDelete)

	err = http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Debug("Server failed")
	}
}
