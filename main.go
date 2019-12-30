package main

import (
	"cake.com/m/config"
	"cake.com/m/controller"
	dataSource "cake.com/m/datasource"
	"cake.com/m/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := AppNew()
	//应用App设置
	configation(app)

	mvcHandler(app)

	config := config.InitConfig()
	port := ":" + config.Port
	app.Run(iris.Addr(port),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, //对json数据序列化更快的配置
	)
}

func mvcHandler(app *iris.Application) {
	engine := dataSource.NewMySqlEngine()

	/**
	用户地址功能
	*/
	addressService := service.NewAddressService(engine)
	address := mvc.New(app.Party("/v1/addr"))
	address.Register(
		addressService,
	)
	address.Handle(new(controller.AddressController))
}

func AppNew() *iris.Application {
	app := iris.New()
	//设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")
	//注册静态资源
	app.StaticWeb("/static", "./static")

	//注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})
	return app
}

func configation(app *iris.Application) {
	//配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    " not found",
			"data":   iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    "500",
			"data":   iris.Map{},
		})
	})
}
