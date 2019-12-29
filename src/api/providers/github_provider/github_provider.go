package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/ajrmzcs/golang-microservices/src/api/clients/restclient"
	"github.com/ajrmzcs/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	// Different types of error should be covered
	// Communication Error
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create new repo in github: %v", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          err.Error(),
		}
	}

	// Read api response body
	bytes, err := ioutil.ReadAll(response.Body)
	// Can't read response body
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          "Invalid response body",
			DocumentationUrl: "",
			Errors:           nil,
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode:       http.StatusInternalServerError,
				Message:          "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %v", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          "error when trying to unmarshall create repo successful response",
		}
	}

	return &result, nil
}