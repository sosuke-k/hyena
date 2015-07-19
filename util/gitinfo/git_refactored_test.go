package gitinfo

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)
import (
	"github.com/bradfitz/iter"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/sosuke-k/hyena/util/git"
)

var showCommitText = `commit fc8e7351b01144fa299c23357a1b4b64708684a4
Author: sosuke-k <snoopies.drum@gmail.com>
Date:   Mon Jun 29 02:53:10 2015 +0900

    hyena auto git commit

diff --git a/chrome.json b/chrome.json
index 81bfe46..079f8b2 100644
--- a/chrome.json
+++ b/chrome.json
@@ -1,7 +1,6 @@
 {
   "0": [
-    "https://github.com/sosuke-k/hyena",
-    "http://ejje.weblio.jp/content/hyena"
+    "https://github.com/sosuke-k/hyena/compare/sosuke-k/git?expand=1"
   ],
   "1": [
     "http://aidiary.hatenablog.com/entry/20110514/1305377659",`

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

var secondCommitLog = `commit 65d8e78f9275df5abebcb94284b1fd49418af517
Author: sosuke-k <snoopies.drum@gmail.com>
Date:   Mon Jun 29 02:47:15 2015 +0900

    hyena auto git commit

diff --git a/atom.json b/atom.json
index 0f7a69f..0acae8f 100644
--- a/atom.json
+++ b/atom.json
@@ -1,5 +1,5 @@
 {
   "0": [
-    "/Users/katososuke/go/src/github.com/sosuke-k/hyena"
+    null
   ]
 }
\ No newline at end of file`

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
		shas := GetSHAArray(projectDir)
		var commits []Commit
		for _, sha := range shas {
			commits = append(commits, *NewCommit(git.Show(projectDir, sha)))
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
		// spew.Dump(commits)
	})
}
