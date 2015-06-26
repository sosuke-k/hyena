package kobito

import (
	"os"
	"path"

	"github.com/sosuke-k/hyena/util/jxa"
)

// Save saves kobito windows info to the config file at disPath
func Save(disPath string) {
	if jxa.Check("com.qiita.Kobito") {
		execJXA("save", disPath)
	}
}

// Restore restores kobito windows by the config file at disPath
func Restore(disPath string) {
	execJXA("restore", disPath)
}

/*
 * the parent directory of disPath must exist
 */
func execJXA(cmd string, disPath string) {
	srcDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/app/kobito")
	fileName := "kobito_" + cmd + "_app.applescript"
	args := []string{"-l", "JavaScript", fileName, disPath}
	jxa.Execute(srcDir, args)
}
