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

func Log(dir string) string {
	return execute(dir, []string{"log"})
}

// Diff return git diff oldCommit newCommit in dir
func Diff(dir string, oldCommit string, newCommit string) string {
	return execute(dir, []string{"diff", oldCommit, newCommit})
}

func execute(dir string, args []string) string {
	return sh.Execute(dir, "git", args)
}
