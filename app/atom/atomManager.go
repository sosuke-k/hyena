package atom

import (
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
	jxa.Execute(srcDir, args)
}
