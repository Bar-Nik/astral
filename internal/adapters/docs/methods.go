package docs

import (
	"context"

	"github.com/gofrs/uuid"
)

// UploadFile implements app.FileStore.
func (c *Client) UploadDocs(ctx context.Context, f app.Avatar) (uuid.UUID, error) {
	id := uuid.Must(uuid.NewV4())

	return id, nil
}

// DownloadFile implements app.FileStore.
func (c *Client) DownloadDocs(ctx context.Context, id uuid.UUID) (*app.Avatar, error) {

	return f, nil
}

// DeleteFile implements app.FileStore.
func (c *Client) DeleteDocs(ctx context.Context, id uuid.UUID) error {

	return nil
}

// Close implements io.Closer.
func (*Client) Close() error {
	return nil
}
