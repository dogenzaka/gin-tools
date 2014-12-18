package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEnumValidator(t *testing.T) {

	enums := []string{"hoge", "fuga"}

	Convey("When creating a Enum validator", t, func() {
		v := Enum{enums}
		Convey("give \"hoge\" paramater", func() {
			So(v.Validate("hoge"), ShouldBeTrue)
		})
		Convey("give \"fuga\" paramater", func() {
			So(v.Validate("fuga"), ShouldBeTrue)
		})
		Convey("give \"foo\" paramater", func() {
			So(v.Validate("foo"), ShouldBeFalse)
		})
		Convey("give \"hogefuga\" paramater", func() {
			So(v.Validate("hogefuga"), ShouldBeFalse)
		})
		Convey("give \"empty\" paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

}
