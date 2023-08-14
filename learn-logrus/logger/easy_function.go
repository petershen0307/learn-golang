package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type logContextKeyType string

const (
	logEntryContextKey logContextKeyType = "logrusEntry"
)

// SetLogEntryToContext will set logrus.Entry to context then can use easy function LogWithCtx to get logrus.Entry from context.
func SetLogEntryToContext(ctx context.Context, entry *logrus.Entry) context.Context {
	return context.WithValue(ctx, logEntryContextKey, entry)
}

// LogWithCtx will get logrus.Entry from context, if logrus.Entry not exist in context, will return a new one.
func LogWithCtx(ctx context.Context) *logrus.Entry {
	if logEntry, ok := ctx.Value(logEntryContextKey).(*logrus.Entry); ok {
		return logEntry
	}
	return logrus.NewEntry(logrus.StandardLogger())
}
