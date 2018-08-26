package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
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
var StableWorldCA string

// StableWorldURL is base url for stable.world
var StableWorldURL string

func setupRootCA() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	configDir := path.Join(usr.HomeDir, ".config", "stable.world")
	log.Println("StableWorldCA", StableWorldCA)
	StableWorldCA = path.Join(configDir, "ca.cert")

	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		return err
	}
	resp, err := http.Get(fmt.Sprintf("%s/ca.cert", StableWorldURL))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(StableWorldCA, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	f, err := os.OpenFile("curl.log", os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)

	StableWorldBucket = os.Getenv("STABLE_WORLD_BUCKET")
	StableWorldURL = os.Getenv("STABLE_WORLD_URL")
	if StableWorldURL == "" {
		StableWorldURL = "http://localhost:3011"
	}
	err = setupRootCA()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Args", os.Args)
	log.Println("StableWorldBucket", StableWorldBucket)
	log.Println("StableWorldURL", StableWorldURL)
	log.Println("StableWorldCA", StableWorldCA)

}
