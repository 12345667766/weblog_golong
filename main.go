package main

import (
	"ms-go-blog/common"
	"ms-go-blog/config"
	"ms-go-blog/logger"
	"ms-go-blog/server"
	"os"
)

func init() {
	cwd, _ := os.Getwd()
	logger.SetRollingDaily(cwd, "log.txt")
	//模板加载
	common.LoadTemplate()
}
func main() {
	//程序入口，一个项目 只能有一个入口
	//web程序，http协议 ip port
	logger.Info("hello, world")
	logger.Info("%v", config.Cfg.Viewer)
	server.App.Start("127.0.0.1", "8080")
}
