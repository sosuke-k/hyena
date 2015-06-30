package git

import "github.com/sosuke-k/hyena/util/sh"

// Init execute git init in dir
func Init(dir string) {
	execute(dir, []string{"init"})
}

// Commit execute git add . and commit -m msg
func Commit(dir string, msg string, force bool) {
	execute(dir, []string{"add", "."})
	if force {
		execute(dir, []string{"commit", "--allow-empty", "-m", msg})
	} else {
		execute(dir, []string{"commit", "-m", msg})
	}
}

func execute(dir string, args []string) {
	sh.Execute(dir, "git", args)
}
