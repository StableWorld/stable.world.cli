package main

import (
	"testing"

	"github.com/go-test/deep"
)

func TestMakeEnv(t *testing.T) {
	// test stuff here...
	//
	expectedEnv := []string{
		"http_proxy=http://sw:BUCKET@example.com",
		"https_proxy=http://sw:BUCKET@example.com",
		"npm_config_cafile=/path",
	}

	givenEnv := makeEnv("http://example.com", "BUCKET", "/path")
	if diff := deep.Equal(givenEnv, expectedEnv); diff != nil {
		t.Error(diff)
	}

}
