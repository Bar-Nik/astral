package session

import (
	"backentrymiddle/cmd/libs/internal/app"
	"context"
)

type ctxMarker struct{}

// NewContext returns context with slog.Logger.
func NewContext(ctx context.Context, session *app.Session) context.Context {
	return context.WithValue(ctx, ctxMarker{}, session)
}

// FromContext returns slog.Logger from context.
func FromContext(ctx context.Context) *app.Session {
	l, ok := ctx.Value(ctxMarker{}).(*app.Session)
	if !ok {
		return nil
	}

	return l
}
