package gitinfo

import "testing"
import . "github.com/smartystreets/goconvey/convey"

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

	Convey("setCommitIdx", t, func() {

		var conv Converter
		idx := 3
		conv.setCommitIdx(idx)

		Convey("commit index should be idx", func() {
			So(conv.commitIdx, ShouldEqual, idx)
		})
	})

	Convey("applyD", t, func() {

		Convey("This isn't yet implemented", nil)
	})
}
