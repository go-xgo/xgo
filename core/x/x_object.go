package x

import (
	"github.com/go-logr/logr"
	"github.com/go-xgo/xgo/log"
)

func Log() logr.Logger {
	return log.DefaultLogger()
}
