package xlogger

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"xyz_golang/internal/config"
)

func TestSetup_Development(t *testing.T) {
	cfg := config.Config{IsDevelopment: true}
	Setup(cfg)
	assert.NotNil(t, Logger)
}

func TestSetup_Production(t *testing.T) {
	cfg := config.Config{}
	Setup(cfg)
	assert.NotNil(t, Logger)
}
