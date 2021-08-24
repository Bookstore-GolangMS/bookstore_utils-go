package logger

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestInfo(t *testing.T) {
	infoMessage := "logging info message"
	tags := zap.Field{
		Key:    "anotherMessage-Key",
		String: "anotherMessage-Value",
		Type:   zapcore.StringType,
	}

	Info(infoMessage, tags)
}

func TestError(t *testing.T) {
	infoMessage := "logging error message"
	err := errors.New("logging an error message")
	tags := zap.Field{
		Key:    "anotherMessage-Key",
		String: "anotherMessage-Value",
		Type:   zapcore.StringType,
	}

	Error(infoMessage, err, tags)
}

func TestGetLogger(t *testing.T) {
	log := GetLogger()
	assert.NotNil(t, log)
}
