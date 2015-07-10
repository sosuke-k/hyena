package git

import (
	"fmt"
	"strconv"
)

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

// 要設計
// ConverLogToHistroy convert log to history
func ConverLogToHistroy(fileName string, log LogStruct) Histories {
	var histories []History
	var currents []*History
	commits := log.Commits
	for i, commit := range commits {
		commit = commits[len(commits)-i-1]
		fmt.Println("commit index is " + strconv.Itoa(i))
		diffs := commit.Diff.Diffs
		for _, diff := range diffs {
			// fmt.Println(diff.BeforeFileName)
			if diff.AfterFileName == "b/"+fileName {
				fmt.Println(diff.AfterFileName)
				if diff.BeforeFileName == "/dev/null" {
					senteces := diff.Add.Sentences
					for _, sentence := range senteces {
						h := History{LineString: sentence.S, LineNumberSequence: []int{sentence.N}}
						histories = append(histories, h)
						currents = append(currents, &histories[len(histories)-1])
					}
				} else {
					dSenteces := diff.Delete.Sentences
					for _, sentence := range dSenteces {
						currents[sentence.N-1].LineNumberSequence = append(currents[sentence.N-1].LineNumberSequence, -1)
					}
				}
			}
		}
	}

	return Histories{FileName: fileName, HistoryArray: histories}
}

func deleteCurrent(currents []*History, deleteIdx int) []*History {
	return []*History{}
}

// funcs reverseCommits(oldCommits []CommitStruct) (newCommits []CommitStruct) {
//
// 	return
// }
