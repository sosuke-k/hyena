package gyena

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)
import (
	"github.com/bradfitz/iter"
	. "github.com/smartystreets/goconvey/convey"
)

func TestExtractAuthor(t *testing.T) {

	Convey("Given appropriate commit info", t, func() {

		Convey("It should return name and email", func() {
			So(extractAuthor(showCommitText), ShouldEqual, "sosuke-k <snoopies.drum@gmail.com>")
		})
	})

	Convey("Given not appropriate commit info", t, func() {

		Convey("It should return empty string", func() {
			So(extractAuthor(""), ShouldEqual, "")
		})
	})
}

func TestNewCommit(t *testing.T) {
	Convey("given commit log by 65d8e78f9275df5abebcb94284b1fd49418af517", t, func() {
		Convey("It should return", func() {
			commit := NewCommit(secondCommitLog)
			So(commit.Diffs[0].D.LineNumbers[0], ShouldEqual, 3)
			So(commit.Diffs[0].A.LineNumbers[0], ShouldEqual, 3)
		})
	})
}

func TestDiffInfo(t *testing.T) {

	Convey("contains with n", t, func() {

		Convey("How can i write this test", nil)
	})
}

func TestSortCommit(t *testing.T) {

	Convey("given git_test commit", t, func() {
		projectDir := "/Users/katososuke/.config/hyena/git_test"
		rep := Repository{Dir: projectDir}
		shas := rep.GetSHAArray()
		var commits []Commit
		for _, sha := range shas {
			commits = append(commits, *NewCommit(rep.Show(sha)))
		}
		commits = sortCommit(commits)
		for i := range iter.N(len(commits) - 1) {
			fmt.Println("iter is " + strconv.Itoa(i))
			b := commits[i].Date.Before(commits[i+1].Date)
			So(b, ShouldEqual, true)
		}
		for _, c := range commits {
			fmt.Println(c.Date.Format(time.RFC1123Z))
		}
	})
}

func TestGetSHAArray(t *testing.T) {
	Convey("given lab project directroy", t, func() {
		projectDir := "/Users/katososuke/.config/hyena/lab"
		rep := Repository{Dir: projectDir}
		shas := rep.GetSHAArray()
		So(len(shas), ShouldEqual, 1)
	})
}
