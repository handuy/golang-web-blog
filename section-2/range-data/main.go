package main

import (
	"github.com/kataras/iris/v12"
)

type Post struct {
	Title string
	Description string
	Image string
	Category string
	Author string
	PublishedDate string
}

func main() {
	app := iris.Default()
	tmpl := iris.HTML("./view", ".html")
	app.RegisterView(tmpl)

	app.HandleDir("/static", "./static")

	app.Get("/posts", func(ctx iris.Context) {
		post := []Post{
			Post{
				Title: "The AI magically removes moving objects from videos.",
				Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit",
				Image: "/static/images/img_1.jpg",
				Category: "Politics",
				Author: "Carrol Atkinson",
				PublishedDate: "July 19, 2019",
			},
			Post{
				Title: "Golang is good",
				Description: "I love Golang",
				Image: "/static/images/img_1.jpg",
				Category: "Nature",
				Author: "Rob Pike",
				PublishedDate: "July 20, 2020",
			},
			Post{
				Title: "Learn Java",
				Description: "Java is old",
				Image: "/static/images/img_1.jpg",
				Category: "Java",
				Author: "Oracle",
				PublishedDate: "July 19, 2018",
			},
		}

		ctx.ViewData("post", post)
		ctx.View("posts/index.html")
	})

	app.Listen(":8080")
}
