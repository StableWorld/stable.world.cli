package common

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

// GetExitCode of a run command
func GetExitCode(err error) int {

	if err == nil {
		return 0
	}

	if exitError, ok := err.(*exec.ExitError); ok {
		ws := exitError.Sys().(syscall.WaitStatus)
		return ws.ExitStatus()
	}

	return 1
}

// WereIs returns the full path to an executable
func WereIs(exe string) string {
	osPaths := strings.Split(os.Getenv("PATH"), ":")
	for _, dirname := range osPaths {
		loc := path.Join(dirname, exe)
		if _, err := os.Stat(loc); err == nil {
			return loc
		}
	}
	return ""
}

// StableWorldBucket is the bucket name
var StableWorldBucket string

// StableWorldURL is base url for stable.world
var StableWorldURL string

func init() {
	StableWorldBucket = os.Getenv("STABLE_WORLD_BUCKET")
	StableWorldURL = os.Getenv("STABLE_WORLD_URL")
	if StableWorldURL == "" {
		StableWorldURL = "http://localhost:3011"
	}

}
