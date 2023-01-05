package x

import (
	"github.com/go-logr/logr"
	"github.com/go-xgo/xgo/database/xdb"
	"github.com/go-xgo/xgo/logging/xlog"
	"gorm.io/gorm"
)

func Log() logr.Logger {
	return xlog.DefaultLogger()
}

func DB() *gorm.DB {
	return xdb.DefaultDB()
}
