package atom

import (
	"fmt"
	"os"
	"path"

	"github.com/sosuke-k/hyena/util/jxa"
)

// Save saves atom windows info to the config file at disPath
func Save(disPath string) {
	if jxa.Check("com.github.atom") {
		execJXA("save", disPath)
	}
}

// Restore restores atom windows by the config file at disPath
func Restore(disPath string) {
	execJXA("restore", disPath)
}

/*
 * the parent directory of disPath must exist
 */
func execJXA(cmd string, disPath string) {
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/app/atom")
	fileName := "atom_" + cmd + "_window.applescript"
	args := []string{"-l", "JavaScript", fileName, disPath}
	fmt.Fprintln(os.Stdout, "Executing osascript "+args[2])
	fmt.Fprintln(os.Stdout, "Waiting for the command to finish...")
	jxa.Execute(srcDir, args)
	fmt.Fprintln(os.Stdout, "Finished osascript "+args[2])
}
