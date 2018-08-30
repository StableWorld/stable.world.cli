package common

import (
	"fmt"
	"os"
	"path"
	"strings"
)

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

// GetExecutable finds an executable from envvar or throws error
func GetExecutable(name string) string {
	envName := fmt.Sprintf("STABLE_WORLD_%s", strings.ToUpper(name))
	Exe := os.Getenv(envName)
	if Exe != "" {
		return Exe
	}
	Exe = WereIs(name)
	if Exe != "" {
		return Exe
	}
	fmt.Printf("Could not find %s exe\n", name)
	os.Exit(1)
	return ""
}
