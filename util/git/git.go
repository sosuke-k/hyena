package git

import (
	"bytes"
	"os/exec"
	"unsafe"

	"github.com/sosuke-k/hyena/util/log"
)

// Execute executes osascript with args
func Execute(dir string, args []string) {
	shCmd := "git"

	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to execute " + shCmd)
	for i := range args {
		hyenaLogger.Println("  " + args[i])
	}

	cmd := exec.Command(shCmd, args...)
	cmd.Dir = dir
	stderr, err := cmd.StderrPipe()
	if err != nil {
		hyenaLogger.Fatalln(err)
	}
	if err := cmd.Start(); err != nil {
		hyenaLogger.Fatalln(err)
	}

	errBuf := new(bytes.Buffer)
	errBuf.ReadFrom(stderr)
	errBytes := errBuf.Bytes()
	errString := *(*string)(unsafe.Pointer(&errBytes))
	hyenaLogger.Println("to display " + shCmd + " log:")
	hyenaLogger.Println(errString)

	if err := cmd.Wait(); err != nil {
		hyenaLogger.Printf("Command finished with error: %v", err)
	} else {
		hyenaLogger.Println("Command finished")
	}
}
