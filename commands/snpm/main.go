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
		fmt.Sprintf("http_proxy=%s", proxyURL),
		fmt.Sprintf("https_proxy=%s", proxyURL),
		fmt.Sprintf("npm_config_cafile=%s", CAPath),
	}
}

func main() {
	common.SetupLog()
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

	npm := common.Wrap("npm", os.Args[1:])
	env := makeEnv(url, bucket, ca)
	npm.SetEnv(env)
	exitCode := npm.Exec()
	os.Exit(exitCode)
}
