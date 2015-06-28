package chrome

import (
	"fmt"
	"os"
	"path"

	"github.com/sosuke-k/hyena/util/jxa"
)

// Save saves chrome windows info to the config file at disPath
func Save(disPath string) {
	execJXA("save", disPath)
}

// Restore restores chrome windows by the config file at disPath
func Restore(disPath string) {
	execJXA("restore", disPath)
}

/*
 * the parent directory of disPath must exist
 */
func execJXA(cmd string, disPath string) {
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/app/chrome")
	fileName := "chrome_" + cmd + "_tab.applescript"
	args := []string{"-l", "JavaScript", fileName, disPath}
	fmt.Fprintln(os.Stdout, "to execete osascript "+args[2])
	fmt.Fprintln(os.Stdout, "waiting for osascript command to finish...")
	jxa.Execute(srcDir, args)
	fmt.Fprintln(os.Stdout, "finished osascript "+args[2])
}
