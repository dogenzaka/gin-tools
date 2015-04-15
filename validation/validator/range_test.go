package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRangeValidator(t *testing.T) {

	Convey("Given Range Validator (min=10, max=100)", t, func() {

		r := Range{Min: 10, Max: 100}

		Convey("Given '50', then true", func() {

			So(r.Validate("50"), ShouldBeTrue)

		})

		Convey("Given '10', then true", func() {

			So(r.Validate("10"), ShouldBeTrue)

		})

		Convey("Given '100', then true", func() {

			So(r.Validate("100"), ShouldBeTrue)

		})

		Convey("Given '9', then false", func() {

			So(r.Validate("9"), ShouldBeFalse)

		})

		Convey("Given '101', then false", func() {

			So(r.Validate("101"), ShouldBeFalse)

		})

		Convey("Given '-1', then false", func() {

			So(r.Validate("-1"), ShouldBeFalse)

		})

		Convey("Given 'a', then false", func() {

			So(r.Validate("a"), ShouldBeFalse)

		})

		Convey("Given empty, then true", func() {

			So(r.Validate(""), ShouldBeTrue)

		})

		Convey("Given blank, then true", func() {

			So(r.Validate(" "), ShouldBeTrue)

		})

	})

}
