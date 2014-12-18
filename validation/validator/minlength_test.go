package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinLengthValidator(t *testing.T) {

	threshold := 10

	Convey("When creating a MinLength validator", t, func() {
		v := MinLength{threshold}
		Convey("give 10 lengths paramater", func() {
			So(v.Validate("1234567890"), ShouldBeTrue)
		})
		Convey("give 9 lengths paramater", func() {
			So(v.Validate("123456789"), ShouldBeFalse)
		})
		Convey("give a 'empty' paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

}
