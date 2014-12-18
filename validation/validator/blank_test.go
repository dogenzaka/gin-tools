package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEmptyValidator(t *testing.T) {

	Convey("When creating a empty validator", t, func() {
		v := NotBlank{}
		Convey("give a 'hoge' paramater", func() {
			So(v.Validate("hoge"), ShouldBeTrue)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeFalse)
		})
		Convey("give ' '(blank) paramater", func() {
			So(v.Validate(" "), ShouldBeFalse)
		})
		Convey("give '　'(blank) paramater", func() {
			So(v.Validate("　"), ShouldBeFalse)
		})
	})

}
