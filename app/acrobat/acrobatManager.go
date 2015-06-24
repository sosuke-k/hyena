package acrobat

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/sosuke-k/hyena/util/jxa"
)

// Save saves acrobat windows info to the config file at disPath
func Save(disPath string) {
	if jxa.Check("com.adobe.Acrobat.Pro") {
		execJXA("save", disPath)
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
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena")
	srcPath := path.Join(srcDir, "app/acrobat/acrobat_"+cmd+"_doc.applescript")
	shCmd := "osascript"
	args := []string{"-l", "JavaScript", srcPath, disPath}
	if err := exec.Command(shCmd, args...).Run(); err != nil {
		fmt.Println(shCmd)
		for i := range args {
			fmt.Println("  " + args[i])
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
