package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreating(test *testing.T) {
	// creating a model
	type post struct {
		Title   string
		Content string
	}

	postInstance := post{Title: "Post #1", Content: "Post #1."}
	postBytes, err := json.Marshal(postInstance)
	if err != nil {
		test.Logf("unable to create the request: %s", err)
		test.FailNow()
	}

	const url = "http://localhost:3000/posts"
	body := bytes.NewReader(postBytes)
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		test.Logf("unable to create the request: %s", err)
		test.FailNow()
	}

	request.Header.Set("Content-Type", "application/json")
}
