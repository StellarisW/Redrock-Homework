package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"main/app/router"
	"main/boot"
)

func init() {
	boot.ConfigSetup()
	boot.LogSetup()
	boot.ORMSetup()
	boot.RedisSetup()
}

func main() {
	s := g.Server()
	s.SetIndexFolder(false) // 是否允许列出Server主目录的文件列表（默认为false）
	s.SetIndexFiles([]string{"index.html"})
	s.SetServerRoot(".") // 设置Server的主目录

	router.InitRouter(s)
	//ctx := gctx.New()
	//port, _ := g.Cfg().Get(ctx, "server.address")
	//s.SetPort(port.Int()) // 设置服务器端口
	s.Run()
}
