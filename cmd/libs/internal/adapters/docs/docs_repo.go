package docs

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type (
	// Config provide connection info for message broker.
	Config struct {
		Secure    bool
		Endpoint  string
		AccessKey string
		SecretKey string
	}
	// Client provided data from and to message broker.
	Client struct {
		store *minio.Client
		// m     Metrics
	}
)

func New(cfg Config) (*Client, error) {
	opts := &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.Secure,
	}
	client, err := minio.New(cfg.Endpoint, opts)
	if err != nil {
		return nil, fmt.Errorf("minio.New: %w, opts: %+v", err, cfg)
	}
	return &Client{
		store: client,
	}, nil
}
