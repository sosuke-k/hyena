package acrobat

import (
	"os"
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
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/app/acrobat")
	fileName := "acrobat_" + cmd + "_doc.applescript"
	args := []string{"-l", "JavaScript", fileName, disPath}
	jxa.Execute(srcDir, args)
}
