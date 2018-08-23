package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/StableWorld/stable.world.cli/commands/common"
)

var logger *log.Logger

// Curl is the executable to run
var Curl string

func init() {
	logger = log.New(os.Stderr, "", 0)
	Curl = os.Getenv("STABLE_WORLD_CURL")
	if Curl == "" {
		Curl = common.WereIs("curl")
	}
	if Curl == "" {
		fmt.Println("Could not find curl exe")
		os.Exit(1)
	}
}

func transformArguments(args []string) []string {
	var result []string
	urlPrefix := fmt.Sprintf("%s/cache/-/%s", common.StableWorldURL, common.StableWorldBucket)
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			result = append(result, arg)
		} else {
			result = append(result, fmt.Sprintf("%s/%s", urlPrefix, arg))
		}
	}
	return result
}

func main() {

	if common.StableWorldBucket == "" {
		fmt.Fprint(os.Stderr, "envvar STABLE_WORLD_BUCKET is required to be set")
		os.Exit(1)
	}
	argsToCmd := transformArguments(os.Args[1:])
	logger.Println("os.Args[0]", Curl, argsToCmd)
	logger.Println("STABLE_WORLD_BUCKET", common.StableWorldBucket)
	logger.Println("STABLE_WORLD_URL", common.StableWorldURL)

	cmd := exec.Command(Curl, argsToCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	logger.Printf("Command finished with error: %v", err)

	exitCode := common.GetExitCode(err)
	logger.Printf("Command finished with exitCode: %v", exitCode)
	os.Exit(exitCode)
}
