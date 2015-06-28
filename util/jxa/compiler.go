package jxa

import (
	"bytes"
	"os"
	"os/exec"
	"path"
	"unsafe"

	"github.com/sosuke-k/hyena/util/log"
)

// Compile compile fileIO.js and move it to lib folder
func Compile() {
	hyenaLogger := logger.GetInstance()

	libDir := path.Join(os.Getenv("HOME"), "Library/Script Libraries")
	if err := os.MkdirAll(libDir, 0777); err != nil {
		hyenaLogger.Fatalln(err)
	}
	libPath := path.Join(libDir, "fileIO.scpt") //i.e. "~/Library/Script Libraries/jsonIO.scpt"
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena")
	srcPath := path.Join(srcDir, "util/jxa/fileIO.js")
	shCmd := "osacompile"
	args := []string{"-l", "JavaScript", "-o", libPath, srcPath}

	hyenaLogger.Println("to execete " + shCmd)
	for i := range args {
		hyenaLogger.Println("  " + args[i])
	}

	cmd := exec.Command(shCmd, args...)
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
