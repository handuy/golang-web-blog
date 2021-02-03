package main

import "github.com/kataras/iris/v12"

type Post struct {
	ID int
	Title string
}

var posts = []Post{
	Post{
		ID: 1,
		Title: "Golang",
	},
	Post{
		ID: 2,
		Title: "Java",
	},
}

func main() {
	app := iris.Default()
	app.Get("/posts", func(ctx iris.Context){
		ctx.JSON(posts)
	})
	app.Listen(":8080")
}