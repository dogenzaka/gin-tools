package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumberValidator(t *testing.T) {

	Convey("When creating a positive number validator", t, func() {
		v := UInt{}
		Convey("give a positive paramater", func() {
			So(v.Validate("1"), ShouldBeTrue)
		})
		Convey("give a begative paramater", func() {
			So(v.Validate("-1"), ShouldBeFalse)
		})
		Convey("give a zero paramater", func() {
			So(v.Validate("0"), ShouldBeFalse)
		})
		Convey("give a 'empty' paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
		Convey("give a 'hoge' paramater", func() {
			So(v.Validate("hoge"), ShouldBeFalse)
		})
	})

}
