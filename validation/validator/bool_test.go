package validator


import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBoolValidator(t *testing.T) {

	Convey("When creating a bool validator", t, func() {
		v := Bool{}
		Convey("give a 'true' paramater", func () {
			So(v.Validate("true"), ShouldBeTrue)
		})
		Convey("give a 'false' paramater", func () {
			So(v.Validate("false"), ShouldBeTrue)
		})
		Convey("give a 'hoge' paramater", func () {
			So(v.Validate("hoge"), ShouldBeFalse)
		})
		Convey("give a 'empty' paramater", func () {
			So(v.Validate(""), ShouldBeFalse)
		})
	})

	Convey("When creating a bool if not empty validator", t, func() {
		v := BoolIfNotEmpty{}
		Convey("give a 'true' paramater", func () {
			So(v.Validate("true"), ShouldBeTrue)
		})
		Convey("give a 'false' paramater", func () {
			So(v.Validate("false"), ShouldBeTrue)
		})
		Convey("give a 'hoge' paramater", func () {
			So(v.Validate("hoge"), ShouldBeFalse)
		})
		Convey("give a 'empty' paramater", func () {
			So(v.Validate(""), ShouldBeTrue)
		})
	})
}
