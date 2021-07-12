package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		test.Logf("unable to send the request: %s", err)
		test.FailNow()
	}
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		test.Logf("unable to read the response: %s", err)
		test.FailNow()
	}

	type createdPost struct {
		ID int
	}

	createdPostInstance := createdPost{}
	err = json.Unmarshal(responseBytes, &createdPostInstance)
	if err != nil {
		test.Logf("unable to unmarshal the response: %s", err)
		test.FailNow()
	}

	test.Logf("post ID: %d", createdPostInstance.ID)
}
