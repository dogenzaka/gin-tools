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
			So(v.Validate(""), ShouldBeTrue)
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

	Convey("When creating a Url validator", t, func() {
		v := URL{}
		Convey("give 'http://hoge.co.jp' paramater", func() {
			So(v.Validate("http://hoge.co.jp"), ShouldBeTrue)
		})
		Convey("give 'http://www.hoge.com' paramater", func() {
			So(v.Validate("http://www.hoge.com"), ShouldBeTrue)
		})
		Convey("give 'http://hoge.com' paramater", func() {
			So(v.Validate("http://hoge.com"), ShouldBeTrue)
		})
		Convey("give 'http://hoge.com/hoge/aaa/' paramater", func() {
			So(v.Validate("http://hoge.com/fuga/aaa/"), ShouldBeTrue)
		})
		Convey("give 'http://hoge.com/hoge/aaa/+params' paramater", func() {
			So(v.Validate("http://hoge.com/fuga/aaa?param1=value1&pram2=value2"), ShouldBeTrue)
		})
		Convey("give 'http://aaa' paramater", func() {
			So(v.Validate("http://aaa"), ShouldBeFalse)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

	Convey("When creating a Domain validator", t, func() {
		v := Domain{}
		Convey("give 'hoge.co.jp' paramater", func() {
			So(v.Validate("hoge.co.jp"), ShouldBeTrue)
		})
		Convey("give 'hoge.com' paramater", func() {
			So(v.Validate("hoge.com"), ShouldBeTrue)
		})
		Convey("give 'hogecom' paramater", func() {
			So(v.Validate("hogecom"), ShouldBeFalse)
		})
		Convey("give 'hoge.com/fuga' paramater", func() {
			So(v.Validate("hoge.com/fuga"), ShouldBeFalse)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

	Convey("When creating a Email validator", t, func() {
		v := Email{}
		Convey("give 'suzuki@hoge.co.jp' paramater", func() {
			So(v.Validate("suzuki@hoge.co.jp"), ShouldBeTrue)
		})
		Convey("give 'suzuki@hoge.com' paramater", func() {
			So(v.Validate("suzuki@hoge.com"), ShouldBeTrue)
		})
		Convey("give 'suzuki@hogecom' paramater", func() {
			So(v.Validate("suzuki@hogecom"), ShouldBeFalse)
		})
		Convey("give 'suzu.ki@hoge.com' paramater", func() {
			So(v.Validate("suzu.ki@hoge.com"), ShouldBeTrue)
		})
		Convey("give 'suzu..ki@hoge.com' paramater", func() {
			So(v.Validate("suzu..ki.@hoge.com"), ShouldBeTrue)
		})
		Convey("give 'suzuki.@hoge.com' paramater", func() {
			So(v.Validate("suzuki.@hoge.com"), ShouldBeTrue)
		})
		Convey("give 'suzuki..@hoge.com' paramater", func() {
			So(v.Validate("suzuki..@hoge.com"), ShouldBeTrue)
		})
		Convey("give empty paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})
}
