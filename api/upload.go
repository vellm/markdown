package api

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kataras/iris"
)

// UploadMarkdownFile creates a new static content
func UploadMarkdownFile(ctx *iris.Context) {
	creds := credentials.NewEnvCredentials()
	_, err := creds.Get()
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
	}
	cfg := aws.NewConfig().WithRegion("eu-central-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	file, info, err := ctx.FormFile("markdown")
	if err != nil {
		ctx.EmitError(iris.StatusBadRequest)
		return
	}

	buf := new(bytes.Buffer)
	size, _ := buf.ReadFrom(file)

	path := "/" + info.Filename
	params := &s3.PutObjectInput{
		Bucket:        aws.String("markdown-prod"),
		Key:           aws.String(path),
		Body:          file,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String("html"),
	}

	resp, err := svc.PutObject(params)
	if err != nil {
		fmt.Printf("bad response: %s", err)
	}
	ctx.JSON(iris.StatusOK, resp)
}
