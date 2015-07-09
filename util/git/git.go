package git

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	Diffs []DiffStruct `json:"diffs"`
}

// DiffStruct struct
type DiffStruct struct {
	BeforeFileName string         `json:"before_file"`
	AfterFileName  string         `json:"after_file"`
	Add            DiffInfoStruct `json:"add"`
	Delete         DiffInfoStruct `json:"delete"`
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
		commits[i] = CommitStruct{SHA: shas[i], Author: author, Message: message}
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

// Show return git show sha in dir
func Show(dir string, sha string) string {
	return execute(dir, []string{"show", sha})
}

// ParseCommitDiff parses commit and diff info returned git show with sha
func ParseCommitDiff(commitString string) CommitDiffStruct {
	lines := RegSplit(commitString, `\n`)
	var indexes []int
	for i, s := range lines {
		tmp := RegSplit(s, `^diff[\s]+`)
		if len(tmp) > 1 {
			indexes = append(indexes, i)
		}
	}
	diffs := make([]DiffStruct, len(indexes))
	for i, idx := range indexes {
		var diff DiffStruct
		var deletedInfo DiffInfoStruct
		deletedInfo.Lines = []string{}
		var addedInfo DiffInfoStruct
		addedInfo.Lines = []string{}
		start := idx
		var end int
		if i+1 < len(indexes) {
			end = indexes[i+1] // - 1
		} else {
			end = len(lines) // - 1
		}
		for _, line := range lines[start:end] {
			if devided := RegSplit(line, `^---[\s]+`); len(devided) > 1 {
				diff.BeforeFileName = devided[1]
			} else if devided := RegSplit(line, `^\+\+\+[\s]+`); len(devided) > 1 {
				diff.AfterFileName = devided[1]
			} else if devided := RegSplit(line, `[\s]*@@[\s]*`); len(devided) > 1 {
				infos := RegSplit(devided[1], `[\s]`)
				for _, info := range infos {
					if info[0] == '-' {
						ints := RegSplit(info[1:], `,`)
						if i, err := strconv.Atoi(ints[0]); err != nil {
							fmt.Fprintln(os.Stderr, err.Error())
						} else {
							deletedInfo.Start = i
						}
						if i, err := strconv.Atoi(ints[1]); err != nil {
							fmt.Fprintln(os.Stderr, err.Error())
						} else {
							deletedInfo.Sum = i
						}
					}
					if info[0] == '+' {
						ints := RegSplit(info[1:], `,`)
						if i, err := strconv.Atoi(ints[0]); err != nil {
							fmt.Fprintln(os.Stderr, err.Error())
						} else {
							addedInfo.Start = i
						}
						if i, err := strconv.Atoi(ints[1]); err != nil {
							fmt.Fprintln(os.Stderr, err.Error())
						} else {
							addedInfo.Sum = i
						}
					}
				}
			} else if devided := RegSplit(line, `^-{1,1}`); len(devided) > 1 {
				deletedInfo.Lines = append(deletedInfo.Lines, devided[1])
			} else if devided := RegSplit(line, `^\+{1,1}`); len(devided) > 1 {
				addedInfo.Lines = append(addedInfo.Lines, devided[1])
			}
		}
		diff.Delete = deletedInfo
		diff.Add = addedInfo
		diffs[i] = diff
	}
	return CommitDiffStruct{Diffs: diffs}
}

// Diff return git diff oldCommit newCommit in dir
func Diff(dir string, oldCommit string, newCommit string) string {
	return execute(dir, []string{"diff", oldCommit, newCommit})
}

func execute(dir string, args []string) string {
	return sh.Execute(dir, "git", args)
}
