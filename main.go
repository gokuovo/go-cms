package main

import (
	"github.com/gin-gonic/gin"
	"go-cms/pkg/db"
	"go-cms/route"
)

var ENV string

func main() {

	//加载配置文件
	config := configmap(ENV)
	//初始化引擎
	r := gin.Default()
	//初始化mysql
	db.CreateDB(config)
	//注册路由
	route.LoadRoute(r)

	r.Run(":9999") // 监听并在 0.0.0.0:9999 上启动服务

}

func configmap(env string) (envconfig string) {
	envmap := make(map[string]string)
	envmap["dev"] = "dev"
	envmap["qa"] = "qa"
	envmap["online"] = "online"
	v2, ok := envmap[env]
	if !ok {
		return "dev"
	}
	return v2
}
