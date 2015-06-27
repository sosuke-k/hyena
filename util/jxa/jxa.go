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

// Execute executes osascript with args
func Execute(dir string, args []string) {
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
	if err := cmd.Start(); err != nil {
		hyenaLogger.Fatalln(err)
	}
	fmt.Fprintln(os.Stdout, "Waiting for "+shCmd+" command to finish...")

	errBuf := new(bytes.Buffer)
	errBuf.ReadFrom(stderr)
	errBytes := errBuf.Bytes()
	errString := *(*string)(unsafe.Pointer(&errBytes))
	hyenaLogger.Println("to display " + shCmd + " log:")
	hyenaLogger.Println(errString)

	if err := cmd.Wait(); err != nil {
		hyenaLogger.Printf("Command finished with error: %v", err)
		fmt.Fprintf(os.Stdout, "Command finished with error: %v", err)
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintln(os.Stdout, "Please see ~/.config/hyena/hyena.log")
	} else {
		hyenaLogger.Println("Command finished")
		fmt.Fprintln(os.Stdout, "Command finished")
	}
}
