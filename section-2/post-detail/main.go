package main

import (
	"strconv"

	"github.com/kataras/iris/v12"
)

type Author struct {
	ID          int
	Name        string
	Avatar      string
	Description string
}

type Post struct {
	ID            int
	Title         string
	Description   string
	Content       string
	Image         string
	Category      []string
	Tag           []string
	Author        Author
	PublishedDate string
}

var author = []Author{
	Author{
		ID:          1,
		Name:        "Cristiano Ronaldo",
		Avatar:      "/static/images/person_1.jpg",
		Description: "Golang developer",
	},
	Author{
		ID:          2,
		Name:        "Lionel Messi",
		Avatar:      "/static/images/person_2.jpg",
		Description: "Java guru",
	},
	Author{
		ID:          3,
		Name:        "Kylian Mbappe",
		Avatar:      "/static/images/person_3.jpg",
		Description: "NodeJS lover",
	},
}

var postList = []Post{
	Post{
		ID:            6,
		Title:         "The AI magically removes moving objects from videos.",
		Description:   "Lorem ipsum dolor sit amet, consectetur adipisicing elit",
		Content:       "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
		Image:         "/static/images/img_1.jpg",
		Category:      []string{"Nature"},
		Tag:           []string{"Travel", "Food"},
		Author:        author[0],
		PublishedDate: "July 19, 2019",
	},
	Post{
		ID:            7,
		Title:         "Golang is good",
		Description:   "I love Golang",
		Content:       "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC",
		Image:         "/static/images/img_2.jpg",
		Category:      []string{"Nature"},
		Tag:           []string{"Travel", "Food"},
		Author:        author[1],
		PublishedDate: "July 20, 2020",
	},
	Post{
		ID:            8,
		Title:         "Learn Java",
		Description:   "Java is old",
		Content:       "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout",
		Image:         "/static/images/img_3.jpg",
		Category:      []string{"Nature"},
		Tag:           []string{"Travel", "Food"},
		Author:        author[2],
		PublishedDate: "July 19, 2018",
	},
	Post{
		ID:            9,
		Title:         "NodeJS",
		Description:   "NodeJS Express",
		Content:       "The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters",
		Image:         "/static/images/img_3.jpg",
		Category:      []string{"Nature"},
		Tag:           []string{"Travel", "Food"},
		Author:        author[2],
		PublishedDate: "July 19, 2018",
	},
	Post{
		ID:            10,
		Title:         "Golang concurrency",
		Description:   "Goroutine",
		Content:       "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form",
		Image:         "/static/images/img_4.jpg",
		Category:      []string{"Nature"},
		Tag:           []string{"Travel", "Food"},
		Author:        author[2],
		PublishedDate: "July 19, 2018",
	},
	Post{
		ID:            11,
		Title:         "Golang concurrency",
		Description:   "Goroutine",
		Content:       "The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested",
		Image:         "/static/images/img_1.jpg",
		Category:      []string{"Nature"},
		Tag:           []string{"Travel", "Food"},
		Author:        author[2],
		PublishedDate: "July 19, 2018",
	},
}

func main() {
	app := iris.Default()
	tmpl := iris.HTML("./view", ".html")
	app.RegisterView(tmpl)

	app.HandleDir("/static", "./static")

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("postList", postList)
		ctx.View("posts/index.html")
	})

	app.Get("/posts/{postID}", func(ctx iris.Context) {
		postIDStr := ctx.Params().Get("postID")
		postIDInt, _ := strconv.Atoi(postIDStr)

		var rsp Post
		for _, post := range postList {
			if post.ID == postIDInt {
				rsp = post
			}
		}

		ctx.ViewData("post", rsp)
		ctx.View("posts/detail.html")
	})

	app.Listen(":8080")
}
