package main

import (
	"fmt"
	"os"

	"github.com/StableWorld/stable.world.cli/commands/common"
)

// Exe is the executable to run
var Exe string

func main() {

	npm := common.GetExecutable("pip")
	argsToCmd := os.Args[1:]
	env := []string{
		fmt.Sprintf("PIP_PROXY=%s", common.StableWorldProxyURL),
		fmt.Sprintf("PIP_CERT=%s", common.StableWorldCA),
	}
	exitCode := common.Exec(npm, argsToCmd, env)
	os.Exit(exitCode)
}
