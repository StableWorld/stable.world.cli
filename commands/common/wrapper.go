package common

import (
	"log"
	"os"
	"os/exec"
)

// WrappedExecutable provides a wrapper around a call
type WrappedExecutable interface {
	Exec() int
	SetEnv([]string)
}

type wrapped struct {
	path string
	cmd  *exec.Cmd
}

func (w wrapped) Exec() int {
	err := w.cmd.Run()
	log.Printf("Command finished with error: %v", err)

	exitCode := GetExitCode(err)
	log.Printf("Command exitCode: %v", exitCode)
	return exitCode
}

func (w wrapped) SetEnv(env []string) {
	for _, param := range env {
		log.Println("ENV:", param)
	}
	w.cmd.Env = append(os.Environ(), env...)
}

// Wrap an executable
func Wrap(name string, args []string) WrappedExecutable {

	exe := GetExecutable(name)

	log.Println("Running:", exe, args)
	cmd := exec.Command(exe, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return wrapped{
		path: exe,
		cmd:  cmd,
	}
}

// npm := GetExecutable("curl")
// argsToCmd := os.Args[1:]

// // Exec runs a command
// func Exec(exe string, args []string, env []string) int {
//
// 	log.Println("Adding ENVVARS: ", env)
// 	log.Println("Running:", exe, args)
// 	cmd := exec.Command(exe, args...)
// 	env = append(os.Environ(), env...)
// 	cmd.Env = env
//
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
//
// 	log.Printf("Running command and waiting for it to finish...")
// 	err := cmd.Run()
// 	log.Printf("Command finished with error: %v", err)
//
// 	exitCode := GetExitCode(err)
// 	log.Printf("Command exitCode: %v", exitCode)
// 	return exitCode
// }
