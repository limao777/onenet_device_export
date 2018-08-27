package config

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

var config_cfg *goconfig.ConfigFile

//加载文件无法使用统一logging
func init() {
	var err error
	config_cfg, err = goconfig.LoadConfigFile("app.conf")
	if err != nil {
		fmt.Println("open conf file error", err)
		for{}
	}
	config_cfg.GetValue("", "")
}

func Get(section string, key string) string {
	target, err := config_cfg.GetValue(section, key)
	if err != nil {
		fmt.Println("[err] get config error", section, key)
		return ""
	}
	return target
}
