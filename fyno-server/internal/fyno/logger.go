package fyno

import (
	"os"

	"github.com/rs/zerolog"
)

func NewLogger() (logger zerolog.Logger, err error) {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", 0755)
	}

	runLogFile, err := os.OpenFile(
		"log/info.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)

	if err != nil {
		return logger, err
	}

	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	return zerolog.New(multi).With().Timestamp().Logger(), nil
}
