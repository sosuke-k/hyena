package acrobat

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/sosuke-k/hyena/util/jxa"
)

// Save saves acrobat windows info to the config file at disPath
func Save(disPath string) {
	if jxa.Check("com.adobe.Acrobat.Pro") {
		execJXA("save", disPath)
	} else {
		fmt.Println("Acrobat Pro is not running")
	}
}

// Restore restores acrobat windows by the config file at disPath
func Restore(disPath string) {
	execJXA("restore", disPath)
}

/*
 * the parent directory of disPath must exist
 */
func execJXA(cmd string, disPath string) {
	srcDir, err := filepath.Abs(filepath.Dir(os.Args[0])) // i.e. $GOPATH/src/github.com/sosuke-k/hyena
	if err != nil {
		log.Fatal(err)
	}
	srcPath := path.Join(srcDir, "acrobat/acrobat_"+cmd+"_doc.applescript")
	shCmd := "osascript"
	args := []string{"-l", "JavaScript", srcPath, disPath}
	if err = exec.Command(shCmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
