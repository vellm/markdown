package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/vellm/vellm.io/api"
	"github.com/vellm/vellm.io/delivery"
)

var env = os.Getenv("ENV")

func main() {
	iris.Get("/:site", delivery.GetSite)
	iris.Post("/api/v1/upload", api.UploadMarkdownFile)

	iris.Listen(":" + os.Getenv("PORT"))
}
