package logging

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// AccessLogger is a middleware for logging access info
func AccessLogger(out io.Writer) gin.HandlerFunc {

	if out == nil {
		out = os.Stdout
	}

	m := &sync.Mutex{}

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		go func(ctx *gin.Context) {

			al := accessLog{
				logInfo: generateLogInfo(ctx, start),
			}

			if err := ctx.LastError(); err != nil {
				al.Error = err
			}

			bytes, err := json.Marshal(al)
			if err != nil {
				panic(err)
			}

			bytes = append(bytes, 10)
			m.Lock()
			defer m.Unlock()
			out.Write(bytes)
		}(c.Copy())
	}
}

// ActivityLogger is a middleware for logging user action info
func ActivityLogger(out io.Writer, getUserID func(c *gin.Context) (string, error)) gin.HandlerFunc {

	if out == nil {
		out = os.Stdout
	}

	m := &sync.Mutex{}

	return func(c *gin.Context) {

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

		go func(ctx *gin.Context) {

			al := activityLog{
				logInfo: generateLogInfo(ctx, start),
			}

			// get to UserID
			uid, err := getUserID(ctx)
			if err != nil {
				panic(err)
			}
			al.UserID = uid
			al.RequestBody = b

			bytes, err := json.Marshal(al)
			if err != nil {
				panic(err)
			}

			bytes = append(bytes, 10)
			m.Lock()
			defer m.Unlock()
			out.Write(bytes)
		}(c.Copy())
	}
}
