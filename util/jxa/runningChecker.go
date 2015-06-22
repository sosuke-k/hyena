package jxa

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// Check checks if the application of identifier is running
func Check(identifier string) (isRunning bool) {
	var (
		cmdOut []byte
		err    error
	)
	srcDir, err := filepath.Abs(filepath.Dir(os.Args[0])) // i.e. $GOPATH/src/github.com/sosuke-k/hyena
	if err != nil {
		log.Fatal(err)
	}
	srcPath := path.Join(srcDir, "util/jxa/running_checker.applescript")
	shCmd := "osascript"
	args := []string{"-l", "JavaScript", srcPath, identifier}
	if cmdOut, err = exec.Command(shCmd, args...).Output(); err != nil {
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
