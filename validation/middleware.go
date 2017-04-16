package validation

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidatePathParam validate for path parameter
func ValidatePathParam(name string, vs ...Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Params.ByName(name)
		for _, v := range vs {
			if !v.Validate(p) {
				c.Abort(http.StatusBadRequest)
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    "bad_params",
					"message": "bad params",
					"params":  map[string]string{name: p},
				})
				return
			}
		}
	}
}

// ValidateRequestParam validate for request parameter
func ValidateRequestParam(name string, vs ...Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Request.FormValue(name)
		for _, v := range vs {
			if !v.Validate(p) {
				c.Abort(http.StatusBadRequest)
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    "bad_params",
					"message": "bad params",
					"params":  map[string]string{name: p},
				})
				return
			}
		}
	}
}

// ValidateRequestHeader validate for request header
func ValidateRequestHeader(name string, vs ...Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.Request.Header.Get(name)
		for _, v := range vs {
			if !v.Validate(h) {
				c.Abort(http.StatusBadRequest)
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    "bad_params",
					"message": "bad params",
					"params":  map[string]string{name: h},
				})
				return
			}
		}
	}
}

// ValidateRequestBody validate for request body
func ValidateRequestBody(vs ...Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.Abort(http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    "bad_params",
				"message": "bad params",
				"params":  map[string]string{"body": "failed read body"},
			})
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		for _, v := range vs {
			if !v.Validate(string(b)) {
				c.Abort(http.StatusBadRequest)
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    "bad_params",
					"message": "bad params",
					"params":  map[string]string{"body": string(b)},
				})
				return
			}
		}
	}
}
