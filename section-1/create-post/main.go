package main

import (
	"strconv"
	"strings"

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
	Post{
		ID: 3,
		Title: "Javascript",
	},
}

func main() {
	app := iris.Default()
	
	app.Get("/posts", func(ctx iris.Context){
		// /posts?search=""
		// /posts
		// /posts?search="java"
		keyword := ctx.URLParam("search")
		var response []Post

		if keyword != "" {
			for _, item := range posts {
				if strings.Contains(item.Title, keyword) {
					response = append(response, item)
				}
			}
		} else {
			response = posts
		}

		ctx.JSON(response)
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

	app.Post("/posts", func(ctx iris.Context){
		postTitle := ctx.PostValue("title")
		postID := len(posts) + 1

		newPost := Post{
			ID: postID,
			Title: postTitle,
		}

		posts = append(posts, newPost)

		ctx.JSON("Tạo bài viết thành công")
	})

	app.Listen(":8080")
}