package delivery

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kataras/iris"
)

// GetSite delivers a requested static content
func GetSite(ctx *iris.Context) {
	creds := credentials.NewEnvCredentials()
	_, err := creds.Get()
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
	}
	cfg := aws.NewConfig().WithRegion("eu-central-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	site := ctx.Param("site")

	params := &s3.GetObjectInput{
		Bucket: aws.String("markdown-prod"),
		Key:    aws.String(site + ".md"),
	}
	resp, err := svc.GetObject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()

	ctx.Markdown(iris.StatusOK, s)
}
