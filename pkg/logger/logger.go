package logger

import (
	"github.com/kataras/golog"
)

func SetLevel(l string) {
	golog.SetLevel(l)
}

func Info(s string) {
	golog.Info(s)
}

func Warn(s string) {
	golog.Warn(s)
}

func Error(s string) {
	golog.Error(s)
}

func Debug(s string) {
	golog.Debug(s)
}

func Fatal(s string) {
	golog.Fatal(s)
}