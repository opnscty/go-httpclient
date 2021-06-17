package gohttp

import (
	"net/http"
	"testing"
)

// [Check] - is it better to use constants?

func TestGetRequestHeaders(t *testing.T) {
	t.Run("Check for correct header propagation", func(t *testing.T) {

		// Initializaton
		client := httpClient{}
		commonHeaders := make(http.Header)
		commonHeaders.Set("Content-Type", "application/json")
		commonHeaders.Set("User-Agent", "go-httpclient")
		client.builder.SetHeaders(commonHeaders)

		// Execution
		requestHeaders := make(http.Header)
		requestHeaders.Set("X-Request-ID", "ABC-123")

		finalHeaders := client.getRequestHeaders(requestHeaders)

		// Validation
		if len(finalHeaders) != 3 {
			t.Error("3 headers expected, received", len(finalHeaders))
		}

		if finalHeaders.Get("Content-Type") != "application/json" {
			t.Error("Invalid content type received.")
		}

		if finalHeaders.Get("User-Agent") != "go-httpclient" {
			t.Error("Invalid user agent received.")
		}

		if finalHeaders.Get("X-Request-ID") != "ABC-123" {
			t.Error("Invalid request ID received.")
		}
	})
}

func TestGetRequestBody(t *testing.T) {

	// Initialization
	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		// Execution
		body, err := client.getRequestBody("", nil)

		// Validation
		if err != nil {
			t.Error("No error expected when passing nil body.")
		}

		if body != nil {
			t.Error("No body expected when passing nil body.")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		// Validation
		if err != nil {
			t.Error("No error expected when marshalling slice as json.")
		}

		if string(body) != `["one", "two"]` {
			t.Error(("Invalid json body obtained."))
		}

	})

	// [Todo] - finish test cases
	// t.Run("BodyWithXml", func(t *testing.T) {
	// 	// Execution
	// 	body, err := client.getRequestBody("Content-Type", "application/xml")
	// })

	// t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
	// 	// Execution
	// 	body, err := client.getRequestBody("Content-Type", "application/json")
	// })
}
