package gitinfo

import (
	"fmt"
	"strconv"

	"github.com/sosuke-k/hyena/util/git"
)

// FileHistories struct
type FileHistories map[string]Histories

// Histories struct
type Histories []*History

// Phistories struct
// type Phistories []*History

// History struct
type History struct {
	LineString         string         `json:"line_string"`
	LineNumberSequence NumberSequence `json:"line_number_sequence"`
}

// NumberSequence struct
type NumberSequence []int

func (p *NumberSequence) init(n int) {
	ns := make([]int, n)
	for i := range ns {
		ns[i] = -1
	}
	*p = ns
}

func (p *NumberSequence) push(n int) {
	ns := *p
	ns = append(ns, n)
	*p = ns
}

func (pfh *FileHistories) append(fileName string, sentence string) *History {
	fh := *pfh
	h := History{LineString: sentence, LineNumberSequence: []int{}}
	histories, ok := fh[fileName]
	if ok == false {
		fh = make(map[string]Histories)
		histories = Histories{}
	}
	histories = append(histories, &h)
	fh[fileName] = histories
	*pfh = fh
	return &h
}

// Converter struct
type Converter struct {
	commitIdx int
	current   map[string]Histories
	next      map[string]Histories
}

func (c *Converter) init() {
	c.current = make(map[string]Histories)
	c.next = make(map[string]Histories)
}

func (c *Converter) setCommitIdx(idx int) {
	c.commitIdx = idx
}

func (c *Converter) applyD(diff Diff, p *FileHistories) {
	name := diff.D.FileName
	fh := *p
	if name == "/dev/null" {
		fh[diff.A.FileName] = Histories{}
		p = &fh
		return
	}
	// for _, n := range diff.D.LineNumbers {
	// 	if len(c.current[name]) < n {
	// 		fmt.Println("n is out of current")
	// 	}
	// 	c.current[name][n-1].LineNumberSequence.push(-1)
	// }
}

func (c *Converter) applyA(diff Diff, p *FileHistories) {
	fmt.Println("============")
	name := diff.A.FileName
	fh := *p
	c.next[name] = Histories{}
	addedSum := 0
	lnIdx := 0
	fmt.Println("diff.A.Sentences length is " + strconv.Itoa(len(diff.A.Sentences)))

	if len(c.current[name]) == 0 {
		for i, s := range diff.A.Sentences {
			ln := diff.A.LineNumbers[i]

			h := fh.append(name, s)
			h.LineNumberSequence.init(c.commitIdx)
			h.LineNumberSequence.push(ln)
			c.next[name] = append(c.next[name], h)
		}
	}

	for i := range c.current[name] {
		ln := i + addedSum + 1
		for diff.A.LineNumbers[lnIdx] == ln {
			s := diff.A.Sentences[lnIdx]

			h := fh.append(name, s)
			h.LineNumberSequence.init(c.commitIdx)
			h.LineNumberSequence.push(ln)
			c.next[name] = append(c.next[name], h)

			lnIdx++
			addedSum++
		}
		c.next[name] = append(c.next[name], c.current[name][i])
	}
	fmt.Println("c.next[name] length is " + strconv.Itoa(len(c.next[name])))
	p = &fh
}

func convertCommitsToHistory(commits *[]Commit, histories *FileHistories) {
	fmt.Println("convertCommitsToHistory")
	var conv Converter
	for _, commit := range *commits {
		for _, diff := range commit.Diffs {
			conv.applyD(diff, histories)
			conv.applyA(diff, histories)
		}
	}
}

// Hoge func
func Hoge() {
	projectDir := "/Users/katososuke/.config/hyena/git_test"
	shas := GetSHAArray(projectDir)
	var commits []Commit
	for _, sha := range shas {
		commits = append(commits, *NewCommit(git.Show(projectDir, sha)))
	}
	var histories FileHistories
	convertCommitsToHistory(&commits, &histories)
}
