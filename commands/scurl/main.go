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
		fmt.Sprintf("https_proxy=%s", proxyURL),
		fmt.Sprintf("http_proxy=%s", proxyURL),
		fmt.Sprintf("CURL_CA_BUNDLE=%s", CAPath),
	}
}

func run() int {
	common.SetupLog()
	defer common.TeardownLogFile()

	url := common.URL()
	bucket, err := common.Bucket()
	if err != nil {
		fmt.Print(err)
		return 1
	}
	ca, err := common.CA(url)
	if err != nil {
		fmt.Print(err)
		return 1
	}

	curl := common.Wrap("curl", os.Args[1:])
	env := makeEnv(url, bucket, ca)
	curl.SetEnv(env)
	exitCode := curl.Exec()
	return exitCode

}

func main() {
	os.Exit(run())
}
