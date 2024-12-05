package main

import (
	"backentrymiddle/internal/adapters/docs"
	"backentrymiddle/internal/adapters/repo"
	"backentrymiddle/internal/app"
	"fmt"
	"log"
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		DB       dbConfig   `yaml:"db"`
		Docs     DocsConfig `yaml:"docs_store"`
		LogLevel int        `yaml:"loglevel"`
	}
	dbConfig struct {
		DSN string `yaml:"dsn"`
	}
	DocsConfig struct {
		Secure    bool   `yaml:"secure"`
		Endpoint  string `yaml:"endpoint"`
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
	}
)

func main() {
	cfg := Config{}
	yamlContent, err := os.ReadFile("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlContent, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(cfg.LogLevel),
	}))

	r, err := repo.New(repo.Config{
		DSN: cfg.DB.DSN,
	})
	if err != nil {
		fmt.Errorf("repo.New: %w", err)
	}
	defer func() {
		err := r.Close()
		if err != nil {
			log.Error("close database connection")
		}
	}()

	docs, err := docs.New(docs.Config{
		Secure:    cfg.Docs.Secure,
		Endpoint:  cfg.Docs.Endpoint,
		AccessKey: cfg.Docs.AccessKey,
		SecretKey: cfg.Docs.SecretKey,
	})
	if err != nil {
		fmt.Errorf("docs.New: %w", err)
	}
	module := app.New(r, docs)

	httpAPI := http.New(ctx, module) // ?

	// r := mux.NewRouter()

	// migrator, err := migrate.New(
	// 	"file://migrations",
	// 	systemConfig.DSN)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// rowSQLConn, err := sql.Open("postgres", systemConfig.DSN)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// repo := repo.NewRepository(rowSQLConn)
	// r.Use(api.Logging(log))

	// ourServer := api.Server{
	// 	Datebase: repo,
	// }
	// r.HandleFunc("/api/register", ourServer.Register).Methods(http.MethodPost)
	// r.HandleFunc("/api/auth", ourServer.Auth).Methods(http.MethodPost)
	// r.HandleFunc("/api/docs", ourServer.UploadDocs).Methods(http.MethodPost)
	// r.HandleFunc("/api/docs", ourServer.GetAllDocs).Methods(http.MethodGet)
	// r.HandleFunc("/api/docs/{id}", ourServer.GetDocs).Methods(http.MethodGet)
	// r.HandleFunc("/api/docs/{id}", ourServer.DeleteDocs).Methods(http.MethodDelete)
	// r.HandleFunc("/api/auth/", ourServer.CloseSession).Methods(http.MethodDelete)

	// err = http.ListenAndServe("127.0.0.1:8080", r)
	// if err != nil {
	// 	log.Debug("Server failed")
	// }
}
