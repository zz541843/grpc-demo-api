package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	//"shop-api/user-web/initialize"
)

func NewZapLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./zap.log",
		"stderr",
		//"stdout",
	}
	return cfg.Build()
}
func Log() {
	proConfig := zap.NewDevelopmentConfig()
	proConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	proConfig.Level.SetLevel(zap.DebugLevel)
	logger, _ := proConfig.Build()
	//logger := zap.New(
	//	zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), nil, zap.DebugLevel),
	//)
	defer logger.Sync()
	logger.Debug("asdf")
	logger.Info("err")
	logger.Warn("Warn", zap.String("aa", "bb"))

}
func SugarLog() {
	proConfig := zap.NewDevelopmentConfig()
	proConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	proConfig.Level.SetLevel(zap.DebugLevel)
	logger, _ := proConfig.Build()
	sugarLogger := logger.Sugar()
	sugarLogger.Debugf("Trying to hit GET request for %s", "url")
	sugarLogger.Errorf("Error fetching URL %s : Error = %s", "url", "err")
	sugarLogger.Infof("Success! statusCode = %s for URL %s", "sdf", "url")
}
func main() {
	//Log()
	SugarLog()
	////1. 初始化logger
	//initialize.InitLogger()`
	//
	////2. 初始化配置文件
	//initialize.InitConfig()
	//
	//3. 初始化routers
	//Router := initialize.Routers()
	////4. 初始化翻译
	//if err := initialize.InitTrans("zh"); err != nil {
	//	panic(err)
	//}
	////5. 初始化srv的连接
	//initialize.InitSrvConn()
	//
	//viper.AutomaticEnv()
	////如果是本地开发环境端口号固定，线上环境启动获取端口号
	//debug := viper.GetBool("MXSHOP_DEBUG")
	//if !debug{
	//	port, err := utils.GetFreePort()
	//	if err == nil {
	//		global.ServerConfig.Port = port
	//	}
	//}
	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径
	*/
	//zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	//if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
	//	zap.S().Panic("启动失败:", err.Error())
	//}

	//logger, _ := NewZapLogger()
	////zap.NewDevelopmentConfig()
	//
	//defer logger.Sync() // flushes buffer, if any
	//sugar := logger.Sugar()
	//url := "https://bilibili.com"
	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", url,
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)

}
