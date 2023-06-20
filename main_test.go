package main

import (
	"io"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	HttpHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	if bodyString != "Hello World\n" {
		t.Error("Response body should be 'Hello World'")
	}

	if response.StatusCode != 200 {
		t.Error("Response status should be 200")
	}
}
