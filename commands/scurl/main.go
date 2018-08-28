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

func main() {

	npm := common.GetExecutable("curl")
	argsToCmd := os.Args[1:]
	env := makeEnv(common.StableWorldURL, common.StableWorldBucket, common.StableWorldCA)
	exitCode := common.Exec(npm, argsToCmd, env)
	os.Exit(exitCode)
}
