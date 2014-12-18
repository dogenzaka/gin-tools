package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxNumberValidator(t *testing.T) {

	threshold := 1000

	Convey("When creating a MaxNumber validator", t, func() {
		v := MaxNumber{threshold}
		Convey("give 999 paramater", func() {
			So(v.Validate("999"), ShouldBeTrue)
		})
		Convey("give 1000 paramater", func() {
			So(v.Validate("1000"), ShouldBeTrue)
		})
		Convey("give 1001 paramater", func() {
			So(v.Validate("1001"), ShouldBeFalse)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

}
