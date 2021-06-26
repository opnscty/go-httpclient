package core

import "net/http"

type HttpClient interface {
	Do(request *http.Request) (*Response, error)
}
