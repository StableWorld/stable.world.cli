package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/StableWorld/stable.world.cli/commands/common"
)

var logger *log.Logger

// PIP is the executable to run
var PIP string

func init() {
	logger = log.New(os.Stderr, "", 0)
	PIP = os.Getenv("STABLE_WORLD_PIP")
	if PIP == "" {
		PIP = common.WereIs("pip")
	}
	if PIP == "" {
		log.Fatalf("Could not find pip exe")
	}
}

func main() {

	if common.StableWorldBucket == "" {
		fmt.Fprint(os.Stderr, "envvar STABLE_WORLD_BUCKET is required to be set")
		return
	}
	argsToCmd := os.Args[1:]
	logger.Println("os.Args[0]", PIP, argsToCmd)
	logger.Println("STABLE_WORLD_BUCKET", common.StableWorldBucket)
	logger.Println("STABLE_WORLD_URL", common.StableWorldURL)

	cmd := exec.Command(PIP, argsToCmd...)
	env := os.Environ()

	env = append(env, fmt.Sprintf("PIP_INDEX_URL=%s/cache/-/%s/https://pypi.org/simple",
		common.StableWorldURL,
		common.StableWorldBucket,
	))
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	logger.Printf("Command finished with error: %v", err)

	exitCode := common.GetExitCode(err)
	logger.Printf("Command finished with exitCode: %v", exitCode)
	os.Exit(exitCode)
}
