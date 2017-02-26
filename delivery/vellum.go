package delivery

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/kataras/iris"
	"github.com/shurcooL/github_flavored_markdown"
)

var apiResponse struct {
	Markdown      string `json:"Markdown"`
	LastModified  string `json:"LastModified"`
	ContentLength int    `json:"ContentLength"`
}

// requestContent requests markdown of API and compile it to HTML
func requestContent(user, vellum string) string {
	apiURL := "https://" + os.Getenv("API_ENDPOINT") + "/" + os.Getenv("API_STAGE")
	creds := credentials.NewEnvCredentials()
	signer := v4.NewSigner(creds)
	url := apiURL + "/users/" + user + "/vellums/" + vellum
	client := new(http.Client)
	req, _ := http.NewRequest("GET", url, nil)
	signer.Sign(req, nil, "execute-api", "eu-central-1", time.Now())
	resp, _ := client.Do(req)

	json.NewDecoder(resp.Body).Decode(&apiResponse)
	markdown, _ := base64.StdEncoding.DecodeString(apiResponse.Markdown)
	html := github_flavored_markdown.Markdown(markdown)
	return string(html)
}

// GetVellum delivers a requested static content
func GetVellum(ctx *iris.Context) {
	user := ctx.Param("user")
	vellum := ctx.Param("vellum")

	ctx.HTML(iris.StatusOK, requestContent(user, vellum))
}
