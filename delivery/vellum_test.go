package delivery

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/shurcooL/github_flavored_markdown"
	"github.com/stretchr/testify/assert"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

var (
	APIEndpoint = "api.test.io"
	APIStage    = "v1"
	APIUrl      = fmt.Sprintf("https://%s/%s", APIEndpoint, APIStage)
)

type APIResponse struct {
	Markdown      string
	LastModified  string
	ContentLength int
}

func TestMain(t *testing.T) {
	os.Setenv("API_ENDPOINT", APIEndpoint)
	os.Setenv("API_STAGE", APIStage)
}

func TestRequestContent(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	markdown := `# Test
  Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
  tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
  vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,
  no sea takimata sanctus est Lorem ipsum dolor sit amet.
  `
	markdownBase64 := base64.StdEncoding.EncodeToString([]byte(markdown))
	expectedHTML := string(github_flavored_markdown.Markdown([]byte(markdown)))

	user := "max"
	vellum := "test"

	apiResponse := APIResponse{markdownBase64, "test", 10}
	jsonResponse, _ := httpmock.NewJsonResponder(200, apiResponse)
	httpmock.RegisterResponder("GET", APIUrl+"/users/"+user+"/vellums/"+vellum, jsonResponse)

	html := requestContent(user, vellum)
	assert.Equal(t, expectedHTML, html)
}
