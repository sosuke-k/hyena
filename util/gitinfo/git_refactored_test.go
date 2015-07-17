package gitinfo

import "testing"
import . "github.com/smartystreets/goconvey/convey"

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

		Convey("It should return name and email", func() {
			So(extractAuthor(""), ShouldEqual, "")
		})
	})
}
