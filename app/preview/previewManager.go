package preview

import (
	"fmt"
	"os"
	"path"

	"github.com/sosuke-k/hyena/util/jxa"
)

// Save saves Preview.app windows info to the config file at disPath
func Save(disPath string) {
	if jxa.Check("com.apple.Preview") {
		execJXA("save", disPath)
	}
}

// Restore restores Preview.app windows by the config file at disPath
func Restore(disPath string) {
	execJXA("restore", disPath)
}

/*
 * the parent directory of disPath must exist
 */
func execJXA(cmd string, disPath string) {
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/app/preview")
	fileName := "preview_" + cmd + "_doc.applescript"
	args := []string{"-l", "JavaScript", fileName, disPath}
	fmt.Fprintln(os.Stdout, "Executing osascript "+args[2])
	fmt.Fprintln(os.Stdout, "Waiting for the command to finish...")
	jxa.Execute(srcDir, args)
	fmt.Fprintln(os.Stdout, "Finished osascript "+args[2])
}
