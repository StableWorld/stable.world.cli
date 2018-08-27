package common

import (
	"log"
	"os"
	"os/exec"
)

// Exec runs a command
func Exec(exe string, args []string, env []string) int {

	log.Println("Adding ENVVARS: ", env)
	log.Println("Running:", exe, args)
	cmd := exec.Command(exe, args...)
	env = append(os.Environ(), env...)
	cmd.Env = env

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)

	exitCode := GetExitCode(err)
	log.Printf("Command exitCode: %v", exitCode)
	return exitCode
}
