package jxa

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

// Check checks if the application of identifier is running
func Check(identifier string) (isRunning bool) {
	var (
		cmdOut []byte
		err    error
	)
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena")
	srcPath := path.Join(srcDir, "util/jxa/running_checker.applescript")
	shCmd := "osascript"
	args := []string{"-l", "JavaScript", srcPath, identifier}
	if cmdOut, err = exec.Command(shCmd, args...).Output(); err != nil {
		fmt.Println(shCmd)
		for i := range args {
			fmt.Println("  " + args[i])
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	outString := string(cmdOut)
	if strings.Contains(outString, "true") {
		isRunning = true
	} else {
		isRunning = false
	}
	return
}
