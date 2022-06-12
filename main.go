package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"main/app/router"
	"main/app/ui"
	"main/app/websocket"
	"main/boot"
	"time"
)

func init() {
	boot.LogSetup()
	boot.ORMSetup()
	boot.RedisSetup()
	boot.GRPCSetup()
}

func main() {
	//s := g.Server()
	//s.SetIndexFolder(false) // 是否允许列出Server主目录的文件列表（默认为false）
	//s.SetIndexFiles([]string{"index.html"})
	//s.SetServerRoot(".") // 设置Server的主目录
	go websocket.StartWebSocket()
	router.WebsocketInit()

	opt := gcmd.Scan("1.登录，2.注册 (请输入数字)\n")
	if opt == "1" {
		uid := gcmd.Scan("请输入账号:\n")
		password := gcmd.Scan("请输入密码:\n")
		content := g.Client().Timeout(3*time.Second).PostContent(context.Background(), "chess.stellaris.wang/api/login", g.Map{
			"uid":      uid,
			"password": password,
		})
		j := gjson.New(content)
		if j.Get("success").Bool() == true {
			fmt.Println(j.Get("msg"))
			num := gcmd.Scan("请选择房间号:")
			fmt.Println(num)
			ui.NewGame()

		} else {
			fmt.Println(j.Get("msg"))
		}
	} else {
		uid := gcmd.Scan("请输入账号:\n")
		nickname := gcmd.Scan("请输入昵称:\n")
		password := gcmd.Scan("请输入密码:\n")
		g.Client().Timeout(3*time.Second).PostContent(context.Background(), "chess.stellaris.wang/api/register", g.Map{
			"uid":      uid,
			"nickname": nickname,
			"password": password,
		})
	}
	//
	//router.InitRouter(s)
	////ctx := gctx.New()
	////port, _ := g.Cfg().Get(ctx, "server.address")
	////s.SetPort(port.Int()) // 设置服务器端口
	//s.Run()
}
