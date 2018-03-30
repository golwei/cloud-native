package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unrolled/render"
)

var formatter = render.New(render.Options{IndentJSON: true})

func TestCreateMatch(t *testing.T) {
	body := []byte(`{"key1":"some val"}`)
	client := &http.Client{}
	server := httptest.NewServer(http.HandlerFunc(createMatch(formatter)))
	defer server.Close()
	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("req err: %v", err)
	}
	res, err := client.Do(req)
	if err != nil {
		t.Errorf("res err: %v", err)
	}
	defer res.Body.Close()
	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("readall err: %v", err)
	}
	fmt.Println(string(resbody))
}
