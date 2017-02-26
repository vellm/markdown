package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/vellm/vellm.io/delivery"
)

var env = os.Getenv("ENV")

func main() {
	iris.Get("/:user/:vellum", delivery.GetVellum)

	iris.Listen(":" + os.Getenv("PORT"))
}
