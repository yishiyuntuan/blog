package logger

import (
	"sync"

	"github.com/kataras/golog"
)

var Log *golog.Logger
var once sync.Once

func init() {
	once.Do(func() {
		Log = GetLogger()
	})
}

func GetLogger() *golog.Logger {
	logger := golog.Default
	logger.SetLevel("debug")

	return logger
}
