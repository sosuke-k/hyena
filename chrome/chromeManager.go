package chrome

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func Save(disPath string) {
	execJXA("save", disPath)
}

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
	srcPath := path.Join(srcDir, "chrome/chrome_"+cmd+"_tab.applescript")
	shCmd := "osascript"
	args := []string{"-l", "JavaScript", srcPath, disPath}
	if err = exec.Command(shCmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
