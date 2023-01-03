package x

import (
	"github.com/go-logr/logr"
	"github.com/go-xgo/xgo/database"
	"github.com/go-xgo/xgo/log"
	"gorm.io/gorm"
)

func Log() logr.Logger {
	return log.DefaultLogger()
}

func DB() *gorm.DB {
	return database.DefaultDB()
}
