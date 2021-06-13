package go_httpclient

import (
	"fmt"
	"io/ioutil"

	"github.com/opsoc/go-httpclient/gohttp"
)

var (
	githubHttpClient = gohttp.NewBuilder()
)

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().Build()

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println((response.StatusCode))

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println((response.StatusCode))

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
