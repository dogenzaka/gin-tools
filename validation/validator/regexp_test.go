package validator

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRegExpValidator(t *testing.T) {

	threshold := "^[a-z0-9]*$"

	Convey("When creating a RegExp validator", t, func() {
		v := RegExp{threshold}
		Convey("give \"123\" paramater", func() {
			So(v.Validate("123"), ShouldBeTrue)
		})
		Convey("give \"abcdefgh\" paramater", func() {
			So(v.Validate("abcdefgh"), ShouldBeTrue)
		})
		Convey("give \"abc100\" paramater", func() {
			So(v.Validate("abc100"), ShouldBeTrue)
		})
		Convey("give \"ABC100\" paramater", func() {
			So(v.Validate("ABC100"), ShouldBeFalse)
		})
		Convey("give \"abc_100\" paramater", func() {
			So(v.Validate("abc_100"), ShouldBeFalse)
		})
		Convey("give \"abc-100\" paramater", func() {
			So(v.Validate("abc-100"), ShouldBeFalse)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

	threshold = "^[a-z]{5}$"

	Convey("When creating a RegExp validator", t, func() {
		v := RegExp{threshold}
		Convey("give \"abcde\" paramater", func() {
			So(v.Validate("abcde"), ShouldBeTrue)
		})
		Convey("give \"abcdef\" paramater", func() {
			So(v.Validate("abcdef"), ShouldBeFalse)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeFalse)
		})
	})

	threshold = "^$"

	Convey("When creating a RegExp validator", t, func() {
		v := RegExp{threshold}
		Convey("give \"123\" paramater", func() {
			So(v.Validate("123"), ShouldBeFalse)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

	threshold = ""

	Convey("When creating a RegExp validator", t, func() {
		v := RegExp{threshold}
		Convey("give \"123\" paramater", func() {
			So(v.Validate("123"), ShouldBeTrue)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})
}
