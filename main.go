package main

import (
	"github.com/HenrikFricke/markdown/api"
	"github.com/HenrikFricke/markdown/delivery"
	"github.com/kataras/iris"
)

func main() {
	iris.Get("/:site", delivery.GetSite)
	iris.Post("/api/v1/upload", api.UploadMarkdownFile)
	iris.Listen(":8080")
}
