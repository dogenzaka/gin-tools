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

	Convey("When creating a UnixTime validator", t, func() {
		v := UnixTime{}
		Convey("give '-10000' paramater", func() {
			So(v.Validate("-10000"), ShouldBeFalse)
		})
		Convey("give '-1' paramater", func() {
			So(v.Validate("-1"), ShouldBeFalse)
		})
		Convey("give '0' paramater", func() {
			So(v.Validate("0"), ShouldBeTrue)
		})
		Convey("give '1' paramater", func() {
			So(v.Validate("1"), ShouldBeTrue)
		})
		Convey("give '1392899576' paramater", func() {
			So(v.Validate("1392899576"), ShouldBeTrue)
		})
		Convey("give '1392899576000' paramater", func() {
			So(v.Validate("1392899576000"), ShouldBeTrue)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

}
