package gohttp_mock

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClientMock struct {
}

func (c *httpClientMock) Do(request *http.Request) (*http.Response, error) {

	requestBody, err := request.GetBody()
	if err != nil {
		return nil, err
	}

	defer requestBody.Close()

	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}

	if mock := GetMock(request.Method, request.URL.String(), string(body)); mock != nil {
		response := http.Response{
			StatusCode:    mock.ResponseStatusCode,
			Body:          ioutil.NopCloser(strings.NewReader(mock.ResponseBody)),
			ContentLength: int64(len(mock.ResponseBody)),
		}
		return &response, nil
	}
	response := http.Response{
		StatusCode:    http.StatusInternalServerError,
		Body:          ioutil.NopCloser(strings.NewReader(mock.ResponseBody)),
		ContentLength: int64(len(mock.ResponseBody)),
	}
	return &response, nil
}
