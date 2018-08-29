package common

import (
	"testing"

	"github.com/go-test/deep"
)

func TestMakeProxyUrl(t *testing.T) {

	if diff := deep.Equal(MakeProxyURL("http://hello.com", "B1"), "http://sw:B1@hello.com"); diff != nil {
		t.Error(diff)
	}
}
