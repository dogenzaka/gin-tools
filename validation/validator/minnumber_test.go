package validator

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMinNumberValidator(t *testing.T) {

	threshold := 1000

	Convey("When creating a MinNumber validator", t, func() {
		v := MinNumber{threshold}
		Convey("give 999 paramater", func() {
			So(v.Validate("999"), ShouldBeFalse)
		})
		Convey("give 1000 paramater", func() {
			So(v.Validate("1000"), ShouldBeTrue)
		})
		Convey("give 1001 paramater", func() {
			So(v.Validate("1001"), ShouldBeTrue)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

}
