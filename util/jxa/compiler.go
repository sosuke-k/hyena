package jxa

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

// Compile compile fileIO.js and move it to lib folder
func Compile() {
	libDir := path.Join(os.Getenv("HOME"), "Library/Script Libraries")
	if err := os.MkdirAll(libDir, 0777); err != nil {
		log.Fatal(err)
	}
	libPath := path.Join(libDir, "fileIO.scpt") //i.e. "~/Library/Script Libraries/jsonIO.scpt"
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena")
	srcPath := path.Join(srcDir, "util/jxa/fileIO.js")
	shCmd := "osacompile"
	args := []string{"-l", "JavaScript", "-o", libPath, srcPath}
	cmd := exec.Command(shCmd, args...)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(shCmd)
		for i := range args {
			fmt.Println("  " + args[i])
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
