package conf

import (
	"flag"
	"fmt"
	"gin-vben-admin/pkg/constant"
	"gin-vben-admin/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Conf struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql  `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Sqlite Sqlite `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Cors   Cors   `mapstructure:"cors" json:"cors" yaml:"cors"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
}

var C *Conf

// Parse 解析配置
func Parse(path ...string) (*viper.Viper, *Conf) {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", utils.RelativePath("config.yaml"), "choose config file.")
		flag.Parse()
		if cpath := os.Getenv(constant.ConfigPath); cpath != "" {
			config = cpath
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&C); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&C); err != nil {
		fmt.Println(err)
	}
	logrus.WithField("config", C).Info("config")
	return v, C
}
