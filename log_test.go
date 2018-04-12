package goutils

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := ZapLogger(zapLogger)
	logger.Sync()
	logger.Info("test", "error", fmt.Errorf("test err"), "string", "abcdefg")
	logger.Set(&DefaultLogInstance{})
	logger.Info("test", "error", fmt.Errorf("test err"), "string", "abcdefg")
}
