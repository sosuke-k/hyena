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

func (pfh *FileHistories) regulate(n int) {
	fh := *pfh
	for k := range fh {
		for i := range fh[k] {
			l := len(fh[k][i].LineNumberSequence)
			if l < n {
				for j := 0; j < n-l; j++ {
					fh[k][i].LineNumberSequence = append(fh[k][i].LineNumberSequence, 0)
				}
			}
		}
	}
	*pfh = fh
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
	fmt.Println("====applyD====")
	fmt.Println("C " + strconv.Itoa(c.commitIdx))
	fmt.Println("D " + diff.D.FileName)
	fmt.Println("A " + diff.A.FileName)
	name := diff.A.FileName

	if name == "b/chrome.json" {
		fmt.Println("current[" + name + "] length is " + strconv.Itoa(len(c.current[name])))
		fmt.Println("D.Sentences length is " + strconv.Itoa(len(diff.D.Sentences)))
	}

	if diff.D.FileName == "/dev/null" {
		fh := *p
		fh[name] = Histories{}
		c.current[name] = Histories{}
		p = &fh
		return
	}

	c.next[name] = Histories{}
	for i := range c.current[name] {
		ln := i + 1
		if diff.D.contains(ln) {
			c.current[name][i].LineNumberSequence.push(-1)
		} else {
			c.next[name] = append(c.next[name], c.current[name][i])
		}
	}
}

func (c *Converter) applyA(diff Diff, p *FileHistories) {
	fmt.Println("====applyA====")
	fmt.Println("C " + strconv.Itoa(c.commitIdx))
	fmt.Println("D " + diff.D.FileName)
	fmt.Println("A " + diff.A.FileName)
	// spew.Dump(diff.A)
	name := diff.A.FileName
	fh := *p
	c.next[name] = Histories{}
	addedSum := 0
	lnIdx := 0

	if len(c.current[name]) == 0 {
		for i, s := range diff.A.Sentences {
			ln := diff.A.LineNumbers[i]

			h := fh.append(name, s)
			h.LineNumberSequence.init(c.commitIdx)
			h.LineNumberSequence.push(ln)
			c.next[name] = append(c.next[name], h)
		}
	}

	if name == "b/chrome.json" {
		fmt.Println("current[" + name + "] length is " + strconv.Itoa(len(c.current[name])))
		fmt.Println("A.Sentences length is " + strconv.Itoa(len(diff.A.Sentences)))
	}

	addedEndFlag := false
	for i := range c.current[name] {
		fmt.Println("current[" + name + "][" + strconv.Itoa(i) + "]")
		// spew.Dump(c.current[name][i])
		ln := i + addedSum + 1
		fmt.Println("ln = " + strconv.Itoa(ln))
		fmt.Println("lnIdx = " + strconv.Itoa(lnIdx))
		for !addedEndFlag && diff.A.LineNumbers[lnIdx] == ln {
			s := diff.A.Sentences[lnIdx]
			fmt.Println("sentence is " + s)

			h := fh.append(name, s)
			h.LineNumberSequence.init(c.commitIdx)
			h.LineNumberSequence.push(ln)

			fmt.Println("====fh.appended====")
			c.next[name] = append(c.next[name], h)

			lnIdx++
			if lnIdx >= len(diff.A.LineNumbers) {
				addedEndFlag = true
			}
			addedSum++
			ln = i + addedSum + 1
			fmt.Println("====lnIdx and addedSum increment====")
		}
		c.current[name][i].LineNumberSequence.push(ln)
		// fmt.Println("====pushed====")
		c.next[name] = append(c.next[name], c.current[name][i])
	}
	p = &fh
}

// ConvertCommitsToHistory func
func ConvertCommitsToHistory(commits []Commit, histories *FileHistories) {
	fmt.Println("====convertCommitsToHistory====")
	var conv Converter
	conv.init()
	commits = sortCommit(commits)
	for i, commit := range commits {
		conv.setCommitIdx(i)
		for _, diff := range commit.Diffs {
			conv.applyD(diff, histories)
			for k := range conv.next {
				conv.current[k] = conv.next[k]
			}
			conv.applyA(diff, histories)
			for k := range conv.next {
				conv.current[k] = conv.next[k]
			}
		}
		histories.regulate(i + 1)
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
	ConvertCommitsToHistory(commits, &histories)
}
