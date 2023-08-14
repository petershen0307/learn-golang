package logger

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_SetContext(t *testing.T) {
	// arrange
	expectedLogEntry := logrus.WithFields(logrus.Fields{
		"traceID": "123456789",
	})
	// act
	ctx := SetLogEntryToContext(context.Background(), expectedLogEntry)
	// assert
	assert.Equal(t, expectedLogEntry, ctx.Value(logEntryContextKey).(*logrus.Entry))
}

func Test_SetContext_ginContext(t *testing.T) {
	// arrange
	expectedLogEntry := logrus.WithFields(logrus.Fields{
		"traceID": "123456789",
	})
	// act
	gctx := SetLogEntryToContext(&gin.Context{}, expectedLogEntry)
	// assert
	assert.Equal(t, expectedLogEntry, gctx.Value(logEntryContextKey).(*logrus.Entry))
}

func Test_LogWithCtx_ContextWithlogrusEntry(t *testing.T) {
	// arrange
	expectedLogEntry := logrus.WithFields(logrus.Fields{
		"traceID": "123456789",
	})
	ctx := SetLogEntryToContext(context.Background(), expectedLogEntry)
	// act
	actualLogEntry := LogWithCtx(ctx)
	// assert
	assert.Equal(t, expectedLogEntry, actualLogEntry)
}

func Test_LogWithCtx_ContextWithoutlogrusEntry(t *testing.T) {
	// arrange
	expectedLogEntry := logrus.NewEntry(logrus.StandardLogger())
	// act
	actualLogEntry := LogWithCtx(context.Background())
	// assert
	assert.Equal(t, expectedLogEntry, actualLogEntry)
}

func Test_LogWithCtx_GinContextWithlogrusEntry(t *testing.T) {
	// arrange
	expectedLogEntry := logrus.WithFields(logrus.Fields{
		"traceID": "123456789",
	})
	gctx := SetLogEntryToContext(&gin.Context{}, expectedLogEntry)
	// act
	actualLogEntry := LogWithCtx(gctx)
	// assert
	assert.Equal(t, expectedLogEntry, actualLogEntry)
}
