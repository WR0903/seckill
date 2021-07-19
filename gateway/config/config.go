package config

import (
	"os"
	conf "pkg/config"

	"github.com/go-kit/kit/log"
	"github.com/spf13/viper"
)

const (
	kConfigType = "CONFIG_TYPE"
)

var Logger log.Logger

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
	viper.AutomaticEnv()
	initDefault()

	if err := conf.LoadRemoteConfig(); err != nil { //这个是从配置服务拉取配置信息，但是如果拉取不成功也没关系，会从本地拉取一次。
		Logger.Log("Fail to load remote config", err)
	}
	if err := conf.Sub("auth", &AuthPermitConfig); err != nil { //从本地拉取信息
		Logger.Log("Fail to parse config", err)
	}
}
func initDefault() {
	viper.SetDefault(kConfigType, "yaml") //设置一下拉取配置的文本类型
}
