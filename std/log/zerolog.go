package main

import (
    "log/slog"
    "os"

    "github.com/rs/zerolog"
    slogzerolog "github.com/samber/slog-zerolog"
)

func main() {
    zerologL := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

    logger := slog.New(
        slogzerolog.Option{Logger: &zerologL}.NewZerologHandler(),
    )

    logger.Info(
        "incoming request",
        slog.String("method", "GET"),
        slog.String("path", "/api/user"),
        slog.Int("status", 200),
    )
}
