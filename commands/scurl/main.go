package main

import (
	"fmt"
	"os"

	"github.com/StableWorld/stable.world.cli/commands/common"
)

// Exe is the executable to run
var Exe string

func main() {

	npm := common.GetExecutable("curl")
	argsToCmd := os.Args[1:]
	env := []string{
		fmt.Sprintf("https_proxy=%s", common.StableWorldProxyURL),
		fmt.Sprintf("http_proxy=%s", common.StableWorldProxyURL),
		fmt.Sprintf("CURL_CA_BUNDLE=%s", common.StableWorldCA),
	}
	exitCode := common.Exec(npm, argsToCmd, env)
	os.Exit(exitCode)
}
