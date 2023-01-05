package xlog

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

type Config struct {
	Path        string `json:"path"`
	File        string `json:"file"`
	Level       string `json:"level"`
	Stdout      bool   `json:"stdout"`
	RotateTime  int    `json:"rotateTime"`
	RotateLimit uint   `json:"rotateLimit"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Path:        DefaultPath,
		File:        DefaultFile,
		Level:       DefaultLevel,
		Stdout:      DefaultStdout,
		RotateTime:  DefaultRotateTime,
		RotateLimit: DefaultRotateLimit,
	}
}

const (
	DefaultFile        = "default.log"
	DefaultPath        = "logs/"
	DefaultStdout      = true
	DefaultLevel       = "debug"
	DefaultRotateTime  = 24
	DefaultRotateLimit = 7
)

var (
	defaultLogger logr.Logger
)

func init() {
	defaultLogger, _ = New(NewDefaultConfig())
}

func Init(cfg *Config) error {
	l, err := New(cfg)
	if err != nil {
		return err
	}
	SetDefaultLogger(l)
	return nil
}

// DefaultLogger returns the default logger.
func DefaultLogger() logr.Logger {
	return defaultLogger
}

func New(cfg *Config) (logr.Logger, error) {
	logger, _ := zap.NewDevelopment()
	return zapr.NewLogger(logger), nil
}

// SetDefaultLogger sets the default logger for package log.
func SetDefaultLogger(l logr.Logger) {
	defaultLogger = l
}
