package main

import (
	"context"

	"github.com/petershen0307/learn-golang/learn-logrus/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
			"traceID":            "traceID",
			"new":                "new",
		},
	})
	ctx := logger.SetLogEntryToContext(context.Background(), logrus.WithFields(logrus.Fields{
		"traceID": "123456789",
	}))
	logger.LogWithCtx(ctx).Info("test trace id")
	log2 := logger.LogWithCtx(ctx).WithField("new", "new1234")
	ctx = logger.SetLogEntryToContext(context.Background(), log2)
	logger.LogWithCtx(ctx).Info("test new")
	log2.Info("test Log2")
}
