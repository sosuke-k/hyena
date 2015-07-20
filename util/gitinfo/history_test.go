package gitinfo

import (
	"fmt"
	"testing"
)

import (
	"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/sosuke-k/hyena/util/git"
)

func TestNumberSequence(t *testing.T) {

	Convey("After NumberSequence called init with n", t, func() {
		n := 3
		var ns NumberSequence
		ns.init(n)

		Convey("ns length should be n", func() {
			So(len(ns), ShouldEqual, n)
		})

		Convey("all element of ns should be -1", func() {
			for _, v := range ns {
				So(v, ShouldEqual, -1)
			}
		})
	})

	Convey("After NumberSequence called push with n", t, func() {
		n := 6
		var ns NumberSequence
		l := 3
		ns.init(l)
		ns.push(n)

		Convey("ns length should be l+1", func() {
			So(len(ns), ShouldEqual, l+1)
		})

		Convey("last element of ns should be n", func() {
			v := ns[len(ns)-1]
			So(v, ShouldEqual, n)
		})
	})
}

func TestFileHistories(t *testing.T) {

	Convey("After FileHistories called append", t, func() {

		var fh FileHistories
		fileName := "b/chrome.json"
		sentence := "  {"
		fh.append(fileName, sentence)

		Convey("fh[fileName] object should exist", func() {
			So(fh[fileName], ShouldNotEqual, nil)
		})

		Convey("last elements' sentence of fh[fileName] should sentence", func() {
			histories := fh[fileName]
			s := histories[len(histories)-1].LineString
			So(s, ShouldEqual, sentence)
		})
	})
}

func TestConverter(t *testing.T) {

	Convey("init", t, func() {
		Convey("not yet", nil)
	})

	Convey("setCommitIdx", t, func() {

		var conv Converter
		idx := 3
		conv.setCommitIdx(idx)

		Convey("commit index should be idx", func() {
			So(conv.commitIdx, ShouldEqual, idx)
		})
	})

	Convey("apply given first commit diff", t, func() {
		projectDir := "/Users/katososuke/.config/hyena/git_test"
		shas := GetSHAArray(projectDir)
		sha := shas[len(shas)-1]
		commit := *NewCommit(git.Show(projectDir, sha))
		diff := commit.Diffs[0]

		var conv Converter
		idx := 0
		conv.init()
		conv.setCommitIdx(idx)
		fh := FileHistories{}

		l1 := len(fh)
		conv.applyD(diff, &fh)
		So(l1, ShouldNotEqual, len(fh))

		l2 := len(fh[diff.A.FileName])
		conv.applyA(diff, &fh)
		So(l2, ShouldNotEqual, len(fh[diff.A.FileName]))

	})
}

func TestConvertCommitsToHistory(t *testing.T) {
	Convey("not yet", t, func() {
		projectDir := "/Users/katososuke/.config/hyena/git_test"
		shas := GetSHAArray(projectDir)
		var commits []Commit
		for _, sha := range shas {
			commits = append(commits, *NewCommit(git.Show(projectDir, sha)))
		}
		fh := FileHistories{}
		ConvertCommitsToHistory(commits, &fh)
		fmt.Println("====spew.Dump(commits)====")
		spew.Dump(commits)
		fmt.Println("====spew.Dump(fh)====")
		spew.Dump(fh)
	})
}
