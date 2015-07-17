package gitinfo

import (
	"fmt"

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

func (c *Converter) init(idx int) {
	c.current = make(map[string]Histories)
	c.next = make(map[string]Histories)
}

func (c *Converter) setCommitIdx(idx int) {
	c.commitIdx = idx
}

func (c *Converter) apply(diff Diff, histories *FileHistories) {

}

func (c *Converter) applyD(diff Diff, histories *FileHistories) {
	name := diff.D.FileName
	if name == "/dev/null" {
		return
	}
	for _, n := range diff.D.LineNumbers {
		if len(c.current[name]) < n {
			fmt.Println("n is out of current")
		}
		c.current[name][n-1].LineNumberSequence.push(-1)
	}
}

func (c *Converter) applyA(diff Diff, histories *FileHistories) {
	name := diff.A.FileName

	// diff.A.Sentences ではんく c.current[name] のループをまわすようにする

	if name == diff.D.FileName {
		// diff[name] の初期化処理
		return
	}

	for i, s := range diff.A.Sentences {
		n := diff.A.LineNumbers[i]

		h := histories.append(name, s)
		h.LineNumberSequence.init(c.commitIdx - 1)
		h.LineNumberSequence.push(n)

		// current1 := current[name][:i]
		// current2 := c.current[name][i+1:]

		c.next[name] = append(c.current[name][:i], h)
		c.next[name] = append(c.current[name], c.current[name][i+1:]...)
	}
	c.current[name] = append(c.current[name])
}

func convertCommitsToHistory(commits *[]Commit, histories *FileHistories) {
	fmt.Println("convertCommitsToHistory")
	var conv Converter
	for _, commit := range *commits {
		for _, diff := range commit.Diffs {
			conv.apply(diff, histories)
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
