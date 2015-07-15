package gitinfo

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/sosuke-k/hyena/util/git"
	"github.com/sosuke-k/hyena/util/re"
)

// Commit struct
type Commit struct {
	SHA     string    `json:"sha"`
	Author  string    `json:"author"`
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
	Diffs   []Diff    `json:"diffs"`
}

// Diff struct
type Diff struct {
	D DiffInfo `json:"deleted_info"`
	A DiffInfo `json:"added_info"`
}

// DiffInfo struct
type DiffInfo struct {
	FileName    string   `json:"file_name"`
	Sentences   []string `json:"sentences"`
	LineNumbers []int    `json:"line_numbers"`
}

// NewCommit initializes Commit struct by sha
func NewCommit(dir string, sha string) (commit *Commit) {
	commit = new(Commit)
	commitString := git.Show(dir, sha)
	// lines := re.Split(commitString, `\n`)
	commit.SHA = extractSHA(commitString)
	commit.Author = extractAuthor(commitString)
	commit.Date = extractDate(commitString)
	commit.Message = extractMessage(commitString)
	return
}

func devideCommit(log string) (diffs []string) {
	reg := regexp.MustCompile(`\ndiff`)
	idxs := reg.FindAllStringIndex(log, -1)
	// diffs = []
	for i := range idxs {
		if i+1 < len(idxs) {
			diffs = append(diffs, log[idxs[i][0]+1:idxs[i+1][0]+1])
		} else {
			diffs = append(diffs, log[idxs[i][0]+1:])
		}

	}
	return
}

func extractSHA(log string) (sha string) {
	sha = re.FindStringSubmatch(log, `^commit\s([a-zA-Z0-9]{40})`)[1]
	return
}

func extractAuthor(log string) (author string) {
	author = re.FindStringSubmatch(log, `Author:\s(.*)`)[1]
	return
}

func extractDate(log string) (t time.Time) {
	date := re.FindStringSubmatch(log, `Date:\s{3}(.*)`)[1]
	ansic := "Mon Jan _2 15:04:05 2006 +0900"
	t, e := time.Parse(ansic, date)
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}
	return
}

func extractMessage(log string) (msg string) {
	res := re.FindString(log, `Date:\s{3}(\n|.)*diff.*`)
	lines := re.Split(res, `\n`)
	msg = ""
	for i := 1; i < len(lines)-2; i++ {
		msg += lines[i] + "\n"
	}
	msg += lines[len(lines)-2]
	return
}
