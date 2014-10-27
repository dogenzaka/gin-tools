package validator


import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFixedLengthValidator(t *testing.T) {

	threshold := 10

	Convey("When creating a FixedLength validator", t, func() {
		v := FixedLength{threshold}
		Convey("give 10 lengths paramater", func () {
			So(v.Validate("1234567890"), ShouldBeTrue)
		})
		Convey("give 9 lengths paramater", func () {
			So(v.Validate("123456789"), ShouldBeFalse)
		})
		Convey("give 11 lengths paramater", func () {
			So(v.Validate("12345678901"), ShouldBeFalse)
		})
		Convey("give a 'empty' paramater", func () {
			So(v.Validate(""), ShouldBeFalse)
		})
	})

	Convey("When creating a FixedLength validator", t, func() {
		v := FixedLengthIfNotEmpty{threshold}
		Convey("give 10 lengths paramater", func () {
			So(v.Validate("1234567890"), ShouldBeTrue)
		})
		Convey("give 9 lengths paramater", func () {
			So(v.Validate("123456789"), ShouldBeFalse)
		})
		Convey("give 11 lengths paramater", func () {
			So(v.Validate("12345678901"), ShouldBeFalse)
		})
		Convey("give a 'empty' paramater", func () {
			So(v.Validate(""), ShouldBeTrue)
		})
	})
}
