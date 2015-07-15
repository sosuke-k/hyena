package gitinfo

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
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

// GetSHAArray returns commit shas of project
func GetSHAArray(projectDir string) []string {
	return extractAllSHA(git.Log(projectDir))
}

// NewCommit initializes Commit struct by sha
func NewCommit(dir string, sha string) (commit *Commit) {
	commit = new(Commit)
	commitString := git.Show(dir, sha)
	commit.SHA = extractSHA(commitString)
	commit.Author = extractAuthor(commitString)
	commit.Date = extractDate(commitString)
	commit.Message = extractMessage(commitString)
	diffs := devideCommit(commitString)
	// fmt.Println(extractStartAndSum(diffs[0]))
	// extractInfo(&diffs[0], commitString)
	for _, diff := range diffs {
		commit.Diffs = append(commit.Diffs, parseDiff(diff))
	}
	spew.Dump(commit)
	return
}

func extractSHA(log string) (sha string) {
	sha = re.FindStringSubmatch(log, `^commit\s([a-zA-Z0-9]{40})`)[1]
	return
}

func extractAllSHA(log string) (shas []string) {
	res := re.FindAllStringSubmatch(log, `commit\s([a-zA-Z0-9]{40})`)
	for _, r := range res {
		shas = append(shas, r[1])
	}
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
	start := re.FindStringIndex(log, `Date:\s{3}.*\n`)[1]
	end := re.FindStringIndex(log, `\ndiff`)[0]
	msg = log[start:end]
	return
}

func devideCommit(log string) (diffs []string) {
	reg := regexp.MustCompile(`\ndiff`)
	idxs := reg.FindAllStringIndex(log, -1)
	for i := range idxs {
		if i+1 < len(idxs) {
			diffs = append(diffs, log[idxs[i][0]+1:idxs[i+1][0]+1])
		} else {
			diffs = append(diffs, log[idxs[i][0]+1:])
		}

	}
	return
}

func parseDiff(log string) (diff Diff) {
	diff.D.FileName = extractFileName(`\-`, log)
	diff.A.FileName = extractFileName(`\+`, log)
	extractInfo(&diff, log)
	return
}

func extractFileName(signal string, diff string) string {
	return re.FindStringSubmatch(diff, signal+`{3}\s([a-zA-Z0-9./]*)`)[1]
}

func extractStartAndSum(diff string) []int {
	ints := re.FindStringSubmatch(diff, `@@\s\-([0-9]*),([0-9]*)\s\+([0-9]*),([0-9]*)\s@@`)
	dStart, err := strconv.Atoi(ints[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	dSum, err := strconv.Atoi(ints[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	aStart, err := strconv.Atoi(ints[3])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	aSum, err := strconv.Atoi(ints[4])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	return []int{dStart, dSum, aStart, aSum}
}

func extractInfo(diff *Diff, text string) {
	ints := extractStartAndSum(text)
	startD := ints[0]
	startA := ints[2]
	nDeleted := 0
	lines := re.Split(re.Split(text, `@@.*@@\n`)[1], `\n`)
	// fmt.Println(lines)
	for i, line := range lines {
		if len(line) <= 0 {
			break
		}
		// fmt.Println("debug:" + line[0:1])
		if line[0:1] == "-" {
			diff.D.Sentences = append(diff.D.Sentences, line[1:])
			diff.D.LineNumbers = append(diff.D.LineNumbers, startD+i)
			nDeleted++
		}
		if line[0:1] == "+" {
			diff.A.Sentences = append(diff.A.Sentences, line[1:])
			diff.A.LineNumbers = append(diff.A.LineNumbers, startA+i)
		}
	}
}
