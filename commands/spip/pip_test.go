package main

import (
	"testing"

	"github.com/go-test/deep"
)

func TestMakeEnv(t *testing.T) {
	// test stuff here...
	//
	expectedEnv := []string{
		"PIP_PROXY=http://sw:BUCKET@example.com",
		"PIP_CERT=/path",
	}

	givenEnv := makeEnv("http://example.com", "BUCKET", "/path")
	if diff := deep.Equal(givenEnv, expectedEnv); diff != nil {
		t.Error(diff)
	}

}
