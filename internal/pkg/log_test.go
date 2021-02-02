package pkg

import (
	"testing"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	logger := NewLog()
	defer logger.Sync()
	logger.Info("you", zap.String("haha", "hihi"))
	t.Error()
}
