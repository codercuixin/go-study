package mylog

import (
	"fmt"
	"log/slog"
)

func Infof(logger *slog.Logger, format string, args ...any){
	logger.Info(fmt.Sprintf(format, args...))
}