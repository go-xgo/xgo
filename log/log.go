package log

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
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
	return zapr.NewLogger(buildZapLogger(cfg)), nil
}

// SetDefaultLogger sets the default logger for package log.
func SetDefaultLogger(l logr.Logger) {
	defaultLogger = l
}

func encodeTime(_ time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// Suppress actual time to keep output constant.
	enc.AppendString("TIMESTAMP")
}

func buildZapLogger(cfg *Config) *zap.Logger {
	// zap gets configured to not panic on invalid log calls
	// and to produce simple, deterministic output on stdout.
	zc := zap.NewProductionConfig()
	zc.OutputPaths = []string{"stdout"}
	zc.ErrorOutputPaths = zc.OutputPaths
	zc.DisableStacktrace = true
	zc.DisableCaller = true
	zc.EncoderConfig.EncodeTime = encodeTime
	z, _ := zc.Build()
	return z
}
