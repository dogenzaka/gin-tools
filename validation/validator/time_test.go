package validator

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeValidator(t *testing.T) {

	format := time.RFC3339

	Convey("When creating a Time validator", t, func() {
		v := Time{format}
		Convey("give \"2014-11-10T11:56:00+09:00\" paramater", func() {
			So(v.Validate("2014-11-10T11:56:00+09:00"), ShouldBeTrue)
		})
		Convey("give \"2014-11-10T11:56:00Z\" paramater", func() {
			So(v.Validate("2014-11-10T11:56:00Z"), ShouldBeTrue)
		})
		Convey("give \"2014/11/10 11:56:00\" paramater", func() {
			So(v.Validate("2014/11/10 11:56:00"), ShouldBeFalse)
		})
		Convey("give \"hoge\" paramater", func() {
			So(v.Validate("hoge"), ShouldBeFalse)
		})
		Convey("give \"\" paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

}
