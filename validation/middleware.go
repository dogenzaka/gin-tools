package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidatePathParam ... validate for path parameter
func ValidatePathParam(name string, vs ...Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Params.ByName(name)
		for _, v := range vs {
			if !v.Validate(p) {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
		}
	}
}

// ValidateRequestParam ... validate for request parameter
func ValidateRequestParam(name string, vs ...Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Request.FormValue(name)
		for _, v := range vs {
			if !v.Validate(p) {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
		}
	}
}
