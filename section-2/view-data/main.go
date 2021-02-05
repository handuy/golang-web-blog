package main

import (
	"github.com/kataras/iris/v12"
)

type Post struct {
	Title string
	PublishedDate string
}

func main() {
	app := iris.Default()
	tmpl := iris.HTML("./view", ".html")
	app.RegisterView(tmpl)

	app.HandleDir("/static", "./static")

	app.Get("/posts", func(ctx iris.Context) {
		post := Post{
			Title: "Hello Golang",
			PublishedDate: "04-02-2021",
		}

		ctx.ViewData("post", post)
		ctx.View("posts/index.html")
	})

	app.Listen(":8080")
}
