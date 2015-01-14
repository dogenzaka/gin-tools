package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJSONValidator(t *testing.T) {

	Convey("When creating a JSON validator", t, func() {
		v := JSON{}
		Convey("give a '{}' paramater", func() {
			So(v.Validate("{}"), ShouldBeTrue)
		})
		Convey("give a 'false' paramater", func() {
			So(v.Validate(""), ShouldBeFalse)
		})
		Convey("give a '{\"hoge\": 123}' paramater", func() {
			So(v.Validate("{\"hoge\": 123}"), ShouldBeTrue)
		})
		Convey("give a '{hoge: []}' paramater", func() {
			So(v.Validate("{hoge: []}"), ShouldBeFalse)
		})
	})

}
