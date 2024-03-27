package utils

import (
	"fmt"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/chzyer/logex"
)

var _ logging.Logger = &Logger{}

type Logger struct {
	raw *logex.Logger
}

func NewLogger(raw *logex.Logger) *Logger {
	newRaw := raw.DownLevel(1)
	return &Logger{&newRaw}
}

func (l *Logger) toArgs(tags []any) string {
	var args []string
	for i := 0; i < len(tags); i += 2 {
		args = append(args, fmt.Sprintf("%v=%v", tags[i], tags[i+1]))
	}
	return strings.Join(args, " ")
}

func (l *Logger) Debug(msg string, tags ...any) {
	l.raw.Debug(msg, l.toArgs(tags))
}

func (l *Logger) Info(msg string, tags ...any) {
	l.raw.Info(msg, l.toArgs(tags))
}

func (l *Logger) Warn(msg string, tags ...any) {
	l.raw.Warn(msg, l.toArgs(tags))
}

func (l *Logger) Error(msg string, tags ...any) {
	l.raw.Error(msg, l.toArgs(tags))
}

func (l *Logger) Fatal(msg string, tags ...any) {
	l.raw.Fatal(msg, l.toArgs(tags))
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.raw.Debugf(template, args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.raw.Infof(template, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.raw.Warnf(template, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.raw.Errorf(template, args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.raw.Fatalf(template, args...)
}
