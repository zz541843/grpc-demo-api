package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	jz "github.com/zz541843/go-utils"
	"go.uber.org/zap"
	"shop-api/user-web/global"
)

func GetEnvInfo(env string) interface{} {
	viper.AutomaticEnv()
	return viper.Get(env)
	//刚才设置的环境变量 想要生效 我们必须得重启goland
}
func ReadConfig(v *viper.Viper) {
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
}
func InitConfig() {
	configFileName := ""
	if GetEnvInfo("ygb") == nil {
		configFileName = "user-web/config-production.yaml"
	} else {
		configFileName = "user-web/config-development.yaml"
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	ReadConfig(v)
	jz.PrintStruct(global.ServerConfig)
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		zap.S().Infof("配置文件产生变化：%s", in.Name)
		ReadConfig(v)
		jz.PrintStruct(global.ServerConfig)
	})
	//debug := GetEnvInfo("MXSHOP_DEBUG")
	//configFilePrefix := "config"
	//configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	//if debug {
	//	configFileName = fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
	//}
	//
	//v := viper.New()
	////文件的路径如何设置
	//v.SetConfigFile(configFileName)
	//if err := v.ReadInConfig(); err != nil {
	//	panic(err)
	//}
	////这个对象如何在其他文件中使用 - 全局变量
	//if err := v.Unmarshal(global.NacosConfig); err != nil {
	//	panic(err)
	//}
	//zap.S().Infof("配置信息: &v", global.NacosConfig)

}
