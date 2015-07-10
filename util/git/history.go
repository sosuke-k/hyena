package git

// Histories struct
type Histories struct {
	FileName     string    `json:"name"`
	HistoryArray []History `json:"histories"`
}

// History struct
type History struct {
	LineString         string `json:"line_string"`
	LineNumberSequence []int  `json:"line_number_sequence"`
}

// ConverLogToHistroy convert log to history
func ConverLogToHistroy(fileName string, log LogStruct) Histories {
	var histories []History
	commits := log.Commits
	for _, commit := range commits {
		diffs := commit.Diff.Diffs
		for _, diff := range diffs {
			if diff.AfterFileName == "b/"+fileName {

			}
		}
	}

	return Histories{FileName: fileName, HistoryArray: histories}
}
