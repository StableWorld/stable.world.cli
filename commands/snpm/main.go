package main

import (
	"fmt"
	"os"

	"github.com/StableWorld/stable.world.cli/commands/common"
)

// Exe is the executable to run
var Exe string

func main() {
	common.Defaults()
	npm := common.GetExecutable("npm")
	argsToCmd := os.Args[1:]
	env := []string{
		fmt.Sprintf("http_proxy=%s", common.StableWorldProxyURL),
		fmt.Sprintf("https_proxy=%s", common.StableWorldProxyURL),
		fmt.Sprintf("npm_config_cafile=%s", common.StableWorldCA),
	}
	exitCode := common.Exec(npm, argsToCmd, env)
	os.Exit(exitCode)
}
