package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/go-xgo/xgo/core/x"
)
import "github.com/spf13/viper"

func Load(confPath string) error {
	var err error
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// the "config.yaml" in config/ folder has higher priority,then ""config.yaml"" in root folder
	viper.AddConfigPath("config/")
	viper.AddConfigPath(".")
	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	if confPath != "" {
		x.Log().Info("loading external configuration...")
		viper.SetConfigFile(confPath)
		err = viper.MergeInConfig()
		if err != nil {
			return err
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		x.Log().Info("config file has been changed")
	})
	return nil
}
