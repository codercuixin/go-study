package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/docker/docker/daemon/logger"
)

type LevelHandler struct {
	level slog.Leveler
	handler slog.Handler
}

// Enabled implements Handler.Enabled by reporting whether
// level is at least as large as h's level.
func (h *LevelHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level.Level()
}

// Handle implements Handler.Handle.
func (h *LevelHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.handler.Handle(ctx, r)
}

// WithAttrs implements Handler.WithAttrs.
func (h *LevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewLevelHandler(h.level, h.handler.WithAttrs(attrs))
}

// WithGroup implements Handler.WithGroup.
func (h *LevelHandler) WithGroup(name string) slog.Handler {
	return NewLevelHandler(h.level, h.handler.WithGroup(name))
}

// Handler returns the Handler wrapped by h.
func (h *LevelHandler) Handler() slog.Handler {
	return h.handler
}



func NewLevelHandler(level slog.Leveler, h slog.Handler) *LevelHandler{
	if lh, ok := h.(*LevelHandler); ok{
		h = lh.Handler()
	}
	return &LevelHandler{level, h}
}



func main() {
	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(NewLevelHandler(slog.LevelWarn, th))
	logger.Info("not printed")
	logger.Warn("printed")

}