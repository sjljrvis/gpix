package logger

import (
	"github.com/withmandala/go-log"
	"os"
)

// Logger pointer
var Logger *log.Logger

func init() {
	logger := log.New(os.Stderr).WithColor()
	Logger = logger
}
