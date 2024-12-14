package main

import (
	"backentrymiddle/cmd/libs/internal/adapters/docs"
	"backentrymiddle/cmd/libs/internal/adapters/repo"
	"backentrymiddle/cmd/libs/internal/api"
	"backentrymiddle/cmd/libs/internal/app"
	"backentrymiddle/internal/logger"
	"backentrymiddle/internal/password"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sipki-tech/database/connectors"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		DB       dbConfig   `yaml:"db"`
		Docs     DocsConfig `yaml:"docs_store"`
		LogLevel int        `yaml:"loglevel"`
	}
	dbConfig struct {
		MigrateDir string `yaml:"migrate_dir"`
		Driver     string `yaml:"driver"`
		DSN        string `yaml:"dsn"`
	}
	DocsConfig struct {
		Secure    bool   `yaml:"secure"`
		Endpoint  string `yaml:"endpoint"`
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
	}
)

const version = "v0.1.0"

const exitCode = 2

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
	// appName := filepath.Base(os.Args[0])

	ctxParent := logger.NewContext(context.Background(), log.With(slog.String(logger.Version.String(), version)))
	ctx, cancel := signal.NotifyContext(ctxParent, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	defer cancel()
	go forceShutdown(ctx)

	err = start(ctx, cfg)
	if err != nil {
		log.Error("shutdown",
			slog.String(logger.Error.String(), err.Error()),
		)
		os.Exit(exitCode)
	}
}

func start(ctx context.Context, cfg Config) error {
	log := logger.FromContext(ctx)

	r, err := repo.New(ctx, repo.Config{
		DSN: connectors.Raw{
			Query: cfg.DB.DSN,
		},
		MigrateDir: cfg.DB.MigrateDir,
		Driver:     cfg.DB.Driver,
	})
	if err != nil {
		return fmt.Errorf("repo.New: %w", err)
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
		return fmt.Errorf("docs.New: %w", err)
	}

	ph := password.New()

	module := app.New(r, docs, ph)

	httpAPI := api.New(ctx, module)

	log.Warn("Server started")
	err = http.ListenAndServe("127.0.0.1:8080", httpAPI)
	if err != nil {
		return fmt.Errorf("server error: %w", err)

	}
	return nil

}

func forceShutdown(ctx context.Context) {
	log := logger.FromContext(ctx)
	const shutdownDelay = 15 * time.Second

	<-ctx.Done()
	time.Sleep(shutdownDelay)

	log.Error("failed to graceful shutdown")
	os.Exit(2)
}

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
