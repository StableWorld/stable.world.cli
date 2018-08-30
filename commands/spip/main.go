package main

import (
	"fmt"
	"os"

	"github.com/StableWorld/stable.world.cli/commands/common"
)

// Exe is the executable to run
var Exe string

func makeEnv(URL string, bucket string, CAPath string) []string {
	proxyURL := common.MakeProxyURL(URL, bucket)
	return []string{
		fmt.Sprintf("PIP_PROXY=%s", proxyURL),
		fmt.Sprintf("PIP_CERT=%s", CAPath),
	}
}

func main() {

	url := common.URL()
	bucket, err := common.Bucket()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	ca, err := common.CA(url)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	pip := common.Wrap("pip", os.Args[1:])
	env := makeEnv(url, bucket, ca)
	pip.SetEnv(env)
	exitCode := pip.Exec()
	os.Exit(exitCode)
}
