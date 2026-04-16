package test

import (
	"os"
	"testing"

	"gonio/internal/config"
	"gonio/internal/pkg/logger"
)

func TestMain(m *testing.M) {
	logger.Init(&config.LogConfig{
		Mode:  "dev",
		Level: "debug",
	})
	os.Exit(m.Run())
}
