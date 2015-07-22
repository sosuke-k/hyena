package gyena

import "github.com/sosuke-k/hyena/util/sh"

// Repository is git manager
type Repository struct {
	Dir string
}

// Init execute git init in dir
func (rep Repository) Init() {
	execute(rep.Dir, []string{"init"})
}

// Commit execute git add . and commit -m msg
func (rep Repository) Commit(msg string, force bool) {
	execute(rep.Dir, []string{"add", "."})
	if force {
		execute(rep.Dir, []string{"commit", "--allow-empty", "-m", msg})
	} else {
		execute(rep.Dir, []string{"commit", "-m", msg})
	}
}

// Log Returns git log string
func (rep Repository) Log() string {
	return execute(rep.Dir, []string{"log"})
}

// Show return git show sha in dir
func (rep Repository) Show(sha string) string {
	return execute(rep.Dir, []string{"show", sha})
}

// Diff return git diff oldCommit newCommit in dir
func (rep Repository) Diff(oldCommit string, newCommit string) string {
	return execute(rep.Dir, []string{"diff", oldCommit, newCommit})
}

func execute(dir string, args []string) string {
	return sh.Execute(dir, "git", args)
}
