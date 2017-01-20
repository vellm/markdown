package main

import (
	"os"

	"github.com/iris-contrib/middleware/secure"
	"github.com/kataras/iris"
	"github.com/vellm/vellm/api"
	"github.com/vellm/vellm/delivery"
)

func main() {
	s := secure.New(secure.Options{
		SSLRedirect:     true,
		SSLProxyHeaders: map[string]string{"x-forwarded-proto": "https"},
	})

	iris.Use(s)

	iris.Get("/:site", delivery.GetSite)
	iris.Post("/api/v1/upload", api.UploadMarkdownFile)

	iris.Listen(":" + os.Getenv("PORT"))
}
