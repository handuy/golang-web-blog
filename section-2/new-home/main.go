package main

import (
	"github.com/kataras/iris/v12"
)

type Author struct {
	ID int
	Name string
	Avatar string
	Description string
}

type Post struct {
	ID int
	Title string
	Description string
	Image string
	Category []string
	Tag []string
	Author Author
	PublishedDate string
}

func main() {
	app := iris.Default()
	tmpl := iris.HTML("./view", ".html")
	app.RegisterView(tmpl)

	app.HandleDir("/static", "./static")

	app.Get("/", func(ctx iris.Context) {
		author := []Author{
			Author{
				ID: 1,
				Name: "Cristiano Ronaldo",
				Avatar: "/static/images/person_1.jpg",
				Description: "Golang developer",
			},
			Author{
				ID: 2,
				Name: "Lionel Messi",
				Avatar: "/static/images/person_2.jpg",
				Description: "Java guru",
			},
			Author{
				ID: 3,
				Name: "Kylian Mbappe",
				Avatar: "/static/images/person_3.jpg",
				Description: "NodeJS lover",
			},
		}

		firstPostList := []Post{
			Post{
				ID: 1,
				Title: "The AI magically removes moving objects from videos.",
				Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit",
				Image: "/static/images/img_1.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[0],
				PublishedDate: "July 19, 2019",
			},
			Post{
				ID: 2,
				Title: "Golang is good",
				Description: "I love Golang",
				Image: "/static/images/img_2.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[1],
				PublishedDate: "July 20, 2020",
			},
			Post{
				ID: 3,
				Title: "Learn Java",
				Description: "Java is old",
				Image: "/static/images/img_v_1.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[2],
				PublishedDate: "July 19, 2018",
			},
			Post{
				ID: 4,
				Title: "NodeJS",
				Description: "NodeJS Express",
				Image: "/static/images/img_3.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[2],
				PublishedDate: "July 19, 2018",
			},
			Post{
				ID: 5,
				Title: "Golang concurrency",
				Description: "Goroutine",
				Image: "/static/images/img_4.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[2],
				PublishedDate: "July 19, 2018",
			},
		}

		secondPostList := []Post{
			Post{
				ID: 6,
				Title: "The AI magically removes moving objects from videos.",
				Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit",
				Image: "/static/images/img_1.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[0],
				PublishedDate: "July 19, 2019",
			},
			Post{
				ID: 7,
				Title: "Golang is good",
				Description: "I love Golang",
				Image: "/static/images/img_2.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[1],
				PublishedDate: "July 20, 2020",
			},
			Post{
				ID: 8,
				Title: "Learn Java",
				Description: "Java is old",
				Image: "/static/images/img_3.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[2],
				PublishedDate: "July 19, 2018",
			},
			Post{
				ID: 9,
				Title: "NodeJS",
				Description: "NodeJS Express",
				Image: "/static/images/img_3.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[2],
				PublishedDate: "July 19, 2018",
			},
			Post{
				ID: 10,
				Title: "Golang concurrency",
				Description: "Goroutine",
				Image: "/static/images/img_4.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[2],
				PublishedDate: "July 19, 2018",
			},
			Post{
				ID: 11,
				Title: "Golang concurrency",
				Description: "Goroutine",
				Image: "/static/images/img_1.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[2],
				PublishedDate: "July 19, 2018",
			},
		}

		thirdPostList := []Post{
			Post{
				ID: 12,
				Title: "The AI magically removes moving objects from videos.",
				Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit",
				Image: "/static/images/img_1.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[0],
				PublishedDate: "July 19, 2019",
			},
			Post{
				ID: 13,
				Title: "Golang is good",
				Description: "I love Golang",
				Image: "/static/images/img_2.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[1],
				PublishedDate: "July 20, 2020",
			},
			Post{
				ID: 14,
				Title: "Learn Java",
				Description: "Java is old",
				Image: "/static/images/img_3.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[2],
				PublishedDate: "July 19, 2018",
			},
			Post{
				ID: 15,
				Title: "NodeJS",
				Description: "NodeJS Express",
				Image: "/static/images/img_4.jpg",
				Category: []string{"Nature"},
				Tag: []string{"Travel", "Food"},
				Author: author[2],
				PublishedDate: "July 19, 2018",
			},
		}

		ctx.ViewData("firstPostList", firstPostList)
		ctx.ViewData("secondPostList", secondPostList)
		ctx.ViewData("thirdPostList", thirdPostList)
		ctx.View("posts/index.html")
	})

	app.Listen(":8080")
}
