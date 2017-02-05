package main

import (
	"os"

	"github.com/iris-contrib/middleware/secure"
	"github.com/kataras/iris"
	"github.com/vellm/vellm.io/api"
	"github.com/vellm/vellm.io/delivery"
)

var env = os.Getenv("ENV")

func main() {
	s := secure.New(secure.Options{
		SSLRedirect:     true,
		SSLProxyHeaders: map[string]string{"x-forwarded-proto": "https"},
		IsDevelopment:   env == "development",
	})

	iris.Use(s)

	iris.Get("/:site", delivery.GetSite)
	iris.Post("/api/v1/upload", api.UploadMarkdownFile)

	iris.Listen(":" + os.Getenv("PORT"))
}
