package jxa

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/sosuke-k/hyena/util/log"
	"github.com/sosuke-k/hyena/util/sh"
)

// Check checks if the application of identifier is running
func Check(identifier string) (isRunning bool) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to check whether " + identifier + " is running")
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/util/jxa")
	fileName := "running_checker.applescript"
	args := []string{"-l", "JavaScript", fileName, identifier}
	outString := Execute(srcDir, args)
	if strings.Contains(outString, "true") {
		isRunning = true
		hyenaLogger.Println(identifier + " is running")
	} else {
		isRunning = false
		hyenaLogger.Println(identifier + " is not running")
	}
	return
}

// Execute executes osascript with args
func Execute(dir string, args []string) (outString string) {
	shCmd := "osascript"

	fmt.Fprintln(os.Stdout, "to execete "+shCmd+" "+args[2])
	fmt.Fprintln(os.Stdout, "waiting for "+shCmd+" command to finish...")

	outString = sh.Execute(dir, shCmd, args)

	return outString
}
