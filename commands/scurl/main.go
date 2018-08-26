package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"

	"github.com/StableWorld/stable.world.cli/commands/common"
)


// Curl is the executable to run
var Curl string

func init() {
	Curl = os.Getenv("STABLE_WORLD_CURL")
	if Curl == "" {
		Curl = common.WereIs("curl")
	}
	if Curl == "" {
		fmt.Println("Could not find curl exe")
		os.Exit(1)
	}
}

func main() {

	if common.StableWorldBucket == "" {
		fmt.Fprint(os.Stderr, "envvar STABLE_WORLD_BUCKET is required to be set")
		os.Exit(1)
	}
	argsToCmd := os.Args[1:]
	log.Println("os.Args[0]", Curl, argsToCmd)
	log.Println("STABLE_WORLD_BUCKET", common.StableWorldBucket)
	log.Println("STABLE_WORLD_URL", common.StableWorldURL)

	cmd := exec.Command(Curl, argsToCmd...)
	env := os.Environ()

	parsedURL, err := url.Parse(common.StableWorldURL)
	if err != nil {
		log.Fatal(err)
	}
	parsedURL.User = url.UserPassword("sw", common.StableWorldBucket)

	log.Println(parsedURL)

	httpsProxy := fmt.Sprintf("https_proxy=%s", parsedURL.String())
	httpProxy := fmt.Sprintf("http_proxy=%s", parsedURL.String())
	ca := fmt.Sprintf("CURL_CA_BUNDLE=%s", common.StableWorldCA)

	log.Println("Adding ENVVARS: ", httpProxy, httpsProxy, ca)
	env = append(env, httpsProxy, httpProxy, ca)
	// env = append(env, http_proxy)
	cmd.Env = env

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Running command and waiting for it to finish...")
	err = cmd.Run()
	log.Printf("Command finished with error: %v", err)

	exitCode := common.GetExitCode(err)
	log.Printf("Command finished with exitCode: %v", exitCode)
	os.Exit(exitCode)
}
