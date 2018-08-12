package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

// CURL is the executable to run
var CURL string
var STABLE_WORLD_BUCKET string

func getExitCode(err error) int {

	if err == nil {
		return 0
	}

	if exitError, ok := err.(*exec.ExitError); ok {
		ws := exitError.Sys().(syscall.WaitStatus)
		return ws.ExitStatus()
	}

	return 1
}

func which() string {
	osPaths := strings.Split(os.Getenv("PATH"), ":")
	for _, dirname := range osPaths {
		loc := path.Join(dirname, "curl")
		if _, err := os.Stat(loc); err == nil {
			return loc
		}
	}
	return ""
}

func init() {
	STABLE_WORLD_BUCKET = os.Getenv("STABLE_WORLD_BUCKET")
	CURL = os.Getenv("STABLE_WORLD_CURL")
	if CURL == "" {
		CURL = which()
	}
	if CURL == "" {
		fmt.Println("Could not find curl exe")
		os.Exit(1)
	}
}

func main() {

	argsToCmd := os.Args[1:]
	fmt.Println("os.Args[0]", CURL, argsToCmd)
	fmt.Println("STABLE_WORLD_BUCKET", STABLE_WORLD_BUCKET)

	cmd := exec.Command(CURL, argsToCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)

	exitCode := getExitCode(err)
	log.Printf("Command finished with exitCode: %v", exitCode)
	os.Exit(getExitCode(err))
}
