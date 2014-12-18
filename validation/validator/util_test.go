package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtil(t *testing.T) {

	Convey("When isBlank is called", t, func() {

		Convey(`Given a "hoge"`, func() {
			So(isBlank("hoge"), ShouldBeFalse)
		})

		Convey(`Given a ""`, func() {
			So(isBlank(""), ShouldBeTrue)
		})

		Convey(`Given a " "`, func() {
			So(isBlank(" "), ShouldBeTrue)
		})

		Convey(`Given a "　"`, func() {
			So(isBlank("　"), ShouldBeTrue)
		})

	})

}
