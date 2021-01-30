package main

import (
	"errors"
	"github-muzilong/go_iris_admin/bootstrap/mysql"
	"github-muzilong/go_iris_admin/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func main() {

	// 0、初始化配置和数据库
	config.Initialize()
	mysql.SetupDB()

	// 1.创建iris 实例
	app := iris.Default()

	// 2.设置错误等级
	app.Logger().SetLevel("debug")

	// 3.注册模板

	// 4.设置静态文件目录
	app.HandleDir("/public", "./public")

	// 5.设置异常页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.StopWithError(ctx.GetStatusCode(), errors.New("有个错误"))
	})

	// 6、启用session
	sess := sessions.New(sessions.Config{Cookie: "_session_id", AllowReclaim: true})
	app.Use(sess.Handler())

	// 7.注册控制器

	// 8、启动服务
	app.Run(iris.Addr(":3000"))

}
