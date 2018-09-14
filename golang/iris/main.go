package main
import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
)
func main(){
	app:=iris.New()	
	app.Logger().SetLevel("debug")

	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Iris",
		DocPath:  "apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New()) // <- IMPORTANT, register the middleware.

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET","/",func(ctx iris.Context){
		ctx.HTML("<h1>Welcome Handle GET</h1>")
	})

	app.Get("/ping",func(ctx iris.Context){
		ctx.WriteString("pong")
	})

	app.Get("/hello",func(ctx iris.Context){
		ctx.JSON(iris.Map{"message":"Hello Iris!"})
	})
	app.Run(iris.Addr(":8080"),iris.WithoutServerError(iris.ErrServerClosed))
}