package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/vellm/vellm/api"
	"github.com/vellm/vellm/delivery"
)

func main() {
	port := os.Getenv("PORT")
	iris.Get("/:site", delivery.GetSite)
	iris.Post("/api/v1/upload", api.UploadMarkdownFile)
	iris.Listen(":" + port)
}
