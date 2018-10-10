package main
import (
	"time"
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
)
var markdownContents = []byte(`
## Hello Markdown

this is a sample of Markdown contents

### 前端页面数据接口
#### 登录 SSO 统一权限
- 获取当前登录用户信息
- 获取当前登录用户权限配置
- 退出清除Token
#### 后返规则配置
##### 航司规则
- 航司规则列表分页查询
- 传入ID获取航司规则
`)
func main(){
	app :=iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/",cache.Handler(10*time.Second),writeMarkdown)
	app.Run(iris.Addr(":8080"))
}

func writeMarkdown(ctx iris.Context){
	println("Handler executed.Content refreshed.");
	ctx.Markdown(markdownContents)
}