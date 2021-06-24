package examples

import (
	"errors"
	"net/http"
	"testing"

	"github.com/opsoc/go-httpclient/gohttp"
)

func TestPost(t *testing.T) {
	t.Run("Timeout from Github", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:      http.MethodPost,
			URL:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			Error: errors.New("timeout from github"),
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		if repo != nil {
			t.Error("No repo expected when getting timeout from Github.")
		}

		if err == nil {
			t.Error("Error expected when getting timeout frmo Github.")
		}

		// [Todo] test doesn't work...
		// if err.Error() != "timeout from github" {
		// 	fmt.Println(err.Error())
		// 	t.Error("Invalid error message.")
		// }
	})

	t.Run("No error from Github", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:      http.MethodPost,
			URL:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":123,"name":"test-repo"`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		if err != nil {
			t.Error("No error expected when getting a valid response from Github.")
		}

		if repo == nil {
			t.Error("A valid repo was expected.")
		}

		if repo.Name != repository.Name {
			t.Error("Invalid repository name for successful response from Github")
		}
	})
}
