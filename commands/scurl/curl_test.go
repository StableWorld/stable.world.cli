package main

import (
	"testing"

	"github.com/go-test/deep"
)

func TestSomething(t *testing.T) {
	// test stuff here...
	//
	expectedEnv := []string{"https_proxy=", "http_proxy=", "CURL_CA_BUNDLE=/path"}
	givenEnv := makeEnv("http://example.com", "BUCKET", "/path")
	if diff := deep.Equal(givenEnv, expectedEnv); diff != nil {
		t.Error(diff)
	}

}
