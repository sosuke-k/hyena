package sh

import (
	"bytes"
	"os/exec"
	"unsafe"

	"github.com/sosuke-k/hyena/util/log"
)

// Execute shCmd with args in dir
func Execute(dir string, shCmd string, args []string) (outString string) {

	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to execete " + shCmd)
	for i := range args {
		hyenaLogger.Println("  " + args[i])
	}

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

	outBuf := new(bytes.Buffer)
	outBuf.ReadFrom(stdout)
	outBytes := outBuf.Bytes()
	outString = *(*string)(unsafe.Pointer(&outBytes))
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
	} else {
		hyenaLogger.Println("command finished")
	}

	return outString
}
