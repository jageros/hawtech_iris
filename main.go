package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Age       int    `json:"age"`
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}
		ctx.Writef("(Unexpected) internal server error")
	})
	app.StaticWeb("/static", "./web/static")
	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})
	app.Done(func(ctx iris.Context) {})
	app.Get("/", routePath)
	app.Get("/{path}", routePath)
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}

func routePath(ctx iris.Context) {
	path := ctx.Path()
	fmt.Println(path)
	switch path {
	case "/case":
		ctx.View("case.html")
	case "/about":
		ctx.View("about.html")
	case "/news":
		ctx.View("news.html")
	case "/newDetail":
		ctx.View("newDetail.html")
	case "/product":
		ctx.View("product.html")
	default:
		ctx.View("index.html")
	}
}
