package main

import (
	"strconv"

	"github.com/kataras/iris/v12"
)

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
	app.Get("/posts/{postId}", func(ctx iris.Context){
		postID := ctx.Params().Get("postId")
		postIdInt, _ := strconv.Atoi(postID)

		var response Post
		for _, item := range posts {
			if item.ID == postIdInt {
				response = item
			}
		}

		ctx.JSON(response)
	})
	app.Listen(":8080")
}