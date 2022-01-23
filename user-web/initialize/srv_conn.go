package initialize

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	jz "github.com/zz541843/go-utils"
	"google.golang.org/grpc/credentials/insecure"

	//_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"shop-api/user-web/global"
	"shop-api/user-web/proto"
)

func InitSrvConn() {
	userConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	} else {
		zap.S().Infof("UserServer Conection SuccessFul!")
	}
	userSrvClient := proto.NewUserClient(userConn)
	if err != nil {
		zap.S().Fatal(err.Error())
		return
	}
	global.UserSrvClient = userSrvClient

	list, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		PageNumber: 0,
		PageSize:   2,
	})
	if err != nil {
		return
	}
	jz.PrintStruct(list)
}

func InitSrvConn2() {
	//从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name))
	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
		return
	}

	//拨号连接用户grpc服务器 跨域的问题 - 后端解决 也可以前端来解决
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】",
			"msg", err.Error(),
		)
	}
	//1. 后续的用户服务下线了 2. 改端口了 3. 改ip了 负载均衡来做

	//2. 已经事先创立好了连接，这样后续就不用进行再次tcp的三次握手
	//3. 一个连接多个groutine共用，性能 - 连接池
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
