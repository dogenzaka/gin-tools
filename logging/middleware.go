package logging

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// recoverLoggingFailure is a recover when failed to logging
var recoverLoggingFailure func(c *gin.Context)

// SetRecoverLoggingFailure is a set recoverLoggingFailure
func SetRecoverLoggingFailure(f func(c *gin.Context)) {
	recoverLoggingFailure = f
}

// AccessLogger is a middleware for logging access info
func AccessLogger(out io.Writer) gin.HandlerFunc {

	if out == nil {
		out = os.Stdout
	}

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		if recoverLoggingFailure != nil {
			defer recoverLoggingFailure(c)
		}

		al := AccessLog{
			LogInfo: GenerateLogInfo(c, start),
		}

		if err := c.Err(); err != nil {
			al.Error = err
		}

		bytes, err := json.Marshal(al)
		if err != nil {
			panic(err)
		}

		out.Write(append(bytes, 10))
	}
}

// ActivityLogger is a middleware for logging user action info
func ActivityLogger(out io.Writer, getExtra func(c *gin.Context) (string, error)) gin.HandlerFunc {

	if out == nil {
		out = os.Stdout
	}

	return func(c *gin.Context) {

		// check a request method
		if c.Request.Method == "GET" {
			return
		}

		start := time.Now()

		var s string
		func() {
			if recoverLoggingFailure != nil {
				defer recoverLoggingFailure(c)
			}
			var err error
			s, err = ConvertToMapFromBody(c)
			if err != nil {
				panic(err)
			}
		}()

		c.Next()

		if recoverLoggingFailure != nil {
			defer recoverLoggingFailure(c)
		}

		// check a response status
		if c.Writer.Status() < 200 || c.Writer.Status() > 299 {
			return
		}

		al := ActivityLog{
			LogInfo:     GenerateLogInfo(c, start),
			RequestBody: s,
		}

		// get to Extra
		var err error
		if getExtra != nil {
			al.Extra, err = getExtra(c)
			if err != nil {
				panic(err)
			}
		}

		bytes, err := json.Marshal(al)
		if err != nil {
			panic(err)
		}

		out.Write(append(bytes, 10))
	}
}
