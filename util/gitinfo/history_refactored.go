package gitinfo

import "fmt"

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
	ns := *p
	//n個 -1 で初期化
	*p = ns
}

func (p *NumberSequence) push(n int) {

}

func (pfh *FileHistories) append(fileName string, sentence string) *History {
	fh := *pfh
	h := History{LineString: sentence, LineNumberSequence: []int{}}
	fh[fileName] = append(fh[fileName], &h)
	*pfh = fh
	return &h
}

// Converter struct
type Converter struct {
	commitIdx int
	current   map[string]Histories
	next      map[string]Histories
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
	for i, s := range diff.A.Sentences {
		n := diff.A.LineNumbers[i]
		// sからhistory生成
		lns := []int{n}
		h := History{LineString: s, LineNumberSequence: lns}

		// current1 := current[name][:i]
		current2 := c.current[name][i+1:]

		c.next[name] = append(c.current[name][:i], &h)
		c.current[name] = append(c.current[name], current2...)
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
		commits = append(commits, *NewCommit(projectDir, sha))
	}
	var histories FileHistories
	convertCommitsToHistory(&commits, &histories)
}
