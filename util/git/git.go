package git

import "github.com/sosuke-k/hyena/util/sh"

// Init execute git init in dir
func Init(dir string) {
	execute(dir, []string{"init"})
}

// Commit execute git add . and commit -m msg
func Commit(dir string, msg string) {
	execute(dir, []string{"add", "."})
	execute(dir, []string{"commit", "-m", msg})
}

func Log(dir string) string {
	return execute(dir, []string{"log"})
}

func execute(dir string, args []string) string {
	return sh.Execute(dir, "git", args)
}
