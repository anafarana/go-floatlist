package floatlist

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFloatList(t *testing.T) {
	list := Floatlist([]float64{0.0, 1.0, 2.0, 5.0, 1.0, 3.0, 4.0})
	Convey("Float list", t, func() {
		Convey("mean", func() {
			So(list.Mean(), ShouldEqual, 16.0/7.0)
		})

		Convey("median", func() {
			So(list.Median(), ShouldEqual, 2.0)
		})

		Convey("mode", func() {
			So(list.Mode(), ShouldEqual, 1.0)
		})

		Convey("getCountByValue", func() {
			So(list.GetCountByValue(1.0), ShouldEqual, 2)
			So(list.GetCountByValue(8.0), ShouldEqual, 0)
		})
	})
}
