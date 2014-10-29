package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxLengthValidator(t *testing.T) {

	threshold := 10

	Convey("When creating a MaxLength validator", t, func() {
		v := MaxLength{threshold}
		Convey("give 10 lengths paramater", func() {
			So(v.Validate("1234567890"), ShouldBeTrue)
		})
		Convey("give 11 lengths paramater", func() {
			So(v.Validate("12345678901"), ShouldBeFalse)
		})
		Convey("give a 'empty' paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})
}
