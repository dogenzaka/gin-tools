package logging

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RecoverLoggingFailure is a recover when failed to logging
var RecoverLoggingFailure func()

// AccessLogger is a middleware for logging access info
func AccessLogger(out io.Writer) gin.HandlerFunc {

	if out == nil {
		out = os.Stdout
	}

	m := &sync.Mutex{}

	return func(c *gin.Context) {

		if RecoverLoggingFailure != nil {
			defer RecoverLoggingFailure()
		}

		start := time.Now()

		c.Next()

		al := accessLog{
			logInfo: generateLogInfo(c, start),
		}

		if err := c.LastError(); err != nil {
			al.Error = err
		}

		bytes, err := json.Marshal(al)
		if err != nil {
			panic(err)
		}

		bytes = append(bytes, 10)

		// Async write
		go func() {
			m.Lock()
			defer m.Unlock()
			out.Write(bytes)
		}()
	}
}

// ActivityLogger is a middleware for logging user action info
func ActivityLogger(out io.Writer, getExtra func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {

	if out == nil {
		out = os.Stdout
	}

	m := &sync.Mutex{}

	return func(c *gin.Context) {

		if RecoverLoggingFailure != nil {
			defer RecoverLoggingFailure()
		}

		// check a request method
		if c.Request.Method == "GET" {
			return
		}

		start := time.Now()
		b, err := convertToMapFromBody(c)
		if err != nil {
			panic(err)
		}

		c.Next()

		// check a response status
		if c.Writer.Status() < 200 || c.Writer.Status() > 299 {
			return
		}

		al := activityLog{
			logInfo:     generateLogInfo(c, start),
			RequestBody: b,
		}

		// get to Extra
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

		bytes = append(bytes, 10)

		// Async write
		go func() {
			m.Lock()
			defer m.Unlock()
			out.Write(bytes)
		}()
	}
}
