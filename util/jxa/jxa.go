package jxa

import (
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
func Execute(dir string, args []string) string {
	return sh.Execute(dir, "osascript", args)
}

// Compile compile fileIO.js and move it to lib folder
func Compile() {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to compile hyena jxa library and put it to applescript library folder")

	libDir := path.Join(os.Getenv("HOME"), "Library/Script Libraries")
	if err := os.MkdirAll(libDir, 0777); err != nil {
		hyenaLogger.Fatalln(err)
	}
	libPath := path.Join(libDir, "fileIO.scpt") //i.e. "~/Library/Script Libraries/jsonIO.scpt"
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/util/jxa")
	fileName := "fileIO.js"
	shCmd := "osacompile"
	args := []string{"-l", "JavaScript", "-o", libPath, fileName}

	sh.Execute(srcDir, shCmd, args)

	hyenaLogger.Println("finished compiling and putting it there")
}
