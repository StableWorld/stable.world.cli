package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path"
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

// // StableWorldBucket is the bucket name
// var StableWorldBucket string
//
// // StableWorldCA is the path to the root cert
// var StableWorldCA string
//
// // StableWorldProxyURL is the path incuding auth to the proxy server
// var StableWorldProxyURL string
//
// // StableWorldURL is base url for stable.world
// var StableWorldURL string

// func setupRootCA() error {
// 	usr, err := user.Current()
// 	if err != nil {
// 		return err
// 	}
//
// 	configDir := path.Join(usr.HomeDir, ".config", "stable.world")
// 	log.Println("StableWorldCA", StableWorldCA)
// 	StableWorldCA = path.Join(configDir, "ca.cert")
//
// 	err = os.MkdirAll(configDir, 0700)
// 	if err != nil {
// 		return err
// 	}
// 	resp, err := http.Get(fmt.Sprintf("%s/ca.cert", StableWorldURL))
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}
// 	err = ioutil.WriteFile(StableWorldCA, body, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// MakeProxyURL takes a url and a bucket an generates url with auth
func MakeProxyURL(URL string, bucket string) string {
	parsedURL, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}
	parsedURL.User = url.UserPassword("sw", bucket)
	return parsedURL.String()
}

// SetupLog sets the logger to the correct file
func SetupLog() error {
	f, err := os.OpenFile("curl.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(f)
	log.Println("Args", os.Args)
	return nil
}

func URL() string {
	StableWorldURL := os.Getenv("STABLE_WORLD_URL")
	if StableWorldURL == "" {
		StableWorldURL = "http://localhost:3011"
	}
	log.Println("StableWorldURL", StableWorldURL)
	return StableWorldURL
}

func Bucket() (string, error) {
	StableWorldBucket := os.Getenv("STABLE_WORLD_BUCKET")
	var err error
	if StableWorldBucket == "" {
		err = fmt.Errorf("envvar STABLE_WORLD_BUCKET is required to be set")
	} else {
		log.Println("StableWorldBucket", StableWorldBucket)
	}
	return StableWorldBucket, err
}

func CA(url string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	configDir := path.Join(usr.HomeDir, ".config", "stable.world")
	StableWorldCA := path.Join(configDir, "ca.cert")

	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(fmt.Sprintf("%s/ca.cert", url))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(StableWorldCA, body, 0644)
	if err != nil {
		return "", err
	}
	log.Println("StableWorldCA", StableWorldCA)
	return StableWorldCA, nil
}
