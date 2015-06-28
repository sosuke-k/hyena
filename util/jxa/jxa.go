package jxa

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"unsafe"

	"github.com/sosuke-k/hyena/util/log"
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
func Execute(dir string, args []string) string {
	shCmd := "osascript"

	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to execete " + shCmd)
	for i := range args {
		hyenaLogger.Println("  " + args[i])
	}
	fmt.Fprintln(os.Stdout, "to execete "+shCmd+" "+args[2])

	cmd := exec.Command(shCmd, args...)
	cmd.Dir = dir
	stderr, err := cmd.StderrPipe()
	if err != nil {
		hyenaLogger.Fatalln(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		hyenaLogger.Fatalln(err)
	}
	if err := cmd.Start(); err != nil {
		hyenaLogger.Fatalln(err)
	}
	fmt.Fprintln(os.Stdout, "waiting for "+shCmd+" command to finish...")

	outBuf := new(bytes.Buffer)
	outBuf.ReadFrom(stdout)
	outBytes := outBuf.Bytes()
	outString := *(*string)(unsafe.Pointer(&outBytes))
	hyenaLogger.Println("to display " + shCmd + " stdout:")
	hyenaLogger.Println(outString)

	errBuf := new(bytes.Buffer)
	errBuf.ReadFrom(stderr)
	errBytes := errBuf.Bytes()
	errString := *(*string)(unsafe.Pointer(&errBytes))
	hyenaLogger.Println("to display " + shCmd + " stderr:")
	hyenaLogger.Println(errString)

	if err := cmd.Wait(); err != nil {
		hyenaLogger.Printf("command finished with error: %v", err)
		fmt.Fprintf(os.Stdout, "command finished with error: %v", err)
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintln(os.Stdout, "please see ~/.config/hyena/hyena.log")
	} else {
		hyenaLogger.Println("command finished")
		fmt.Fprintln(os.Stdout, "command finished")
	}

	return outString
}
