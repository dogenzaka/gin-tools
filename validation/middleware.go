package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidatePathParam ... validate for path parameter
func ValidatePathParam(name string, vs ...Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, v := range vs {
			if !v.Validate(c.Params.ByName(name)) {
				c.Abort(http.StatusBadRequest)
				return
			}
		}
	}
}

// ValidatePathParam ... validate for request parameter
func ValidateRequestParam(name string, vs ...Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, v := range vs {
			if !v.Validate(c.Request.FormValue(name)) {
				c.Abort(http.StatusBadRequest)
				return
			}
		}
	}
}
