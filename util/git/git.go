package git

import (
	"regexp"
	"time"

	"github.com/sosuke-k/hyena/util/sh"
)

// LogStruct struct
type LogStruct struct {
	Commits []CommitStruct `json:"commits"`
}

// CommitStruct struct
type CommitStruct struct {
	SHA     string           `json:"sha"`
	Author  string           `json:"author"`
	Date    time.Time        `json:"date"`
	Message string           `json:"message"`
	Diff    CommitDiffStruct `json:"diff"`
}

// CommitDiffStruct struct
type CommitDiffStruct struct {
	FileName string       `json:"file"`
	Diffs    []DiffStruct `json:"diffs"`
}

// DiffStruct struct
type DiffStruct struct {
	Add    DiffInfoStruct `json:"add"`
	Delete DiffInfoStruct `json:"delete"`
}

// DiffInfoStruct struct
type DiffInfoStruct struct {
	Start int      `json:"start"`
	Sum   int      `json:"sum"`
	Lines []string `json:"lines"`
}

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

// ParseLog parses git commit log
func ParseLog(logString string) LogStruct {
	lines := RegSplit(logString, `\n`)
	var indexes []int
	var shas []string
	for i, s := range lines {
		tmp := RegSplit(s, `^commit[\s]+`)
		if len(tmp) > 1 {
			indexes = append(indexes, i)
			shas = append(shas, tmp[1])
		}
	}
	commits := make([]CommitStruct, len(indexes))
	for i, idx := range indexes {
		author := RegSplit(lines[idx+1], `^Author:[\s]+`)[1]
		// date := RegSplit(lines[idx+2], `^Date:[\s]+`)[1]
		start := idx + 3
		var end int
		if i+1 < len(indexes) {
			end = indexes[i+1] // - 1
		} else {
			end = len(lines) // - 1
		}
		message := ""
		for _, s := range lines[start:end] {
			message += s + "\n"
		}
		commit := CommitStruct{SHA: shas[i], Author: author, Message: message}
		commits = append(commits, commit)
	}
	return LogStruct{Commits: commits}
}

// RegSplit split text by delimeter
func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}

// Diff return git diff oldCommit newCommit in dir
func Diff(dir string, oldCommit string, newCommit string) string {
	return execute(dir, []string{"diff", oldCommit, newCommit})
}

func execute(dir string, args []string) string {
	return sh.Execute(dir, "git", args)
}
