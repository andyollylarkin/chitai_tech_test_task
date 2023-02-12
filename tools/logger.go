package tools

import (
	"github.com/rs/zerolog"
	"os"
)

func NewLogger() zerolog.Logger {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return l
}
