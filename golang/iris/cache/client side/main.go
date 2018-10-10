package main
import(
	"time"
	"github.com/kataras/iris"
)

const refreshEvery = 10 * time.Second

func main(){
	app := iris.New()
	app.Use(iris.Cache304(refreshEvery))

	app.Get("/",greet);
	app.Run(iris.Addr(":8080"));

}

func greet(ctx iris.Context){
	ctx.Header("X-Custom","my custom header");
	ctx.Writef("Hello World! %s",time.Now())
}