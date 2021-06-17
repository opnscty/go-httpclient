package gohttp

import (
	"fmt"
	"net/http"
)

// Mock provides a clean way to configure http mocks based on a combination
// of a Request Method, URL, and Request Body.
type Mock struct {
	Method      string
	URL         string
	RequestBody string
	// [Todo] Headers?

	ResponseStatusCode int
	ResponseBody       string
	Error              error
}

// GetResponse returns a Response object based on the Mock configuration.
func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := Response{
		status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		statusCode: m.ResponseStatusCode,
		body:       []byte(m.ResponseBody),
	}

	return &response, nil
}
