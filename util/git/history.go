package gyena

// FileHistories struct
type FileHistories map[string]Histories

// Histories struct
type Histories []*History

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
	name := diff.A.FileName

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

	addedEndFlag := false
	for i := range c.current[name] {
		ln := i + addedSum + 1
		if len(diff.A.LineNumbers) == 0 {
			addedEndFlag = true
		}
		for !addedEndFlag && diff.A.LineNumbers[lnIdx] == ln {
			s := diff.A.Sentences[lnIdx]

			h := fh.append(name, s)
			h.LineNumberSequence.init(c.commitIdx)
			h.LineNumberSequence.push(ln)

			c.next[name] = append(c.next[name], h)

			lnIdx++
			if lnIdx >= len(diff.A.LineNumbers) {
				addedEndFlag = true
			}
			addedSum++
			ln = i + addedSum + 1
		}
		c.current[name][i].LineNumberSequence.push(ln)
		c.next[name] = append(c.next[name], c.current[name][i])
	}
	p = &fh
}

func (c *Converter) update(name string) {
	c.current[name] = make([]*History, len(c.next[name]))
	copy(c.current[name], c.next[name])
}

// ConvertCommitsToHistory func
func ConvertCommitsToHistory(commits []Commit, histories *FileHistories) {
	var conv Converter
	conv.init()
	commits = sortCommit(commits)
	for i, commit := range commits {
		conv.setCommitIdx(i)
		for _, diff := range commit.Diffs {
			name := diff.A.FileName
			conv.applyD(diff, histories)
			conv.update(name)
			conv.applyA(diff, histories)
			conv.update(name)
		}
		histories.regulate(i + 1)
	}
}
