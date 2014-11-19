package logging

import (
	"bytes"
	"io/ioutil"
	"time"

	"encoding/json"
	"github.com/gin-gonic/gin"
)

func GenerateLogInfo(c *gin.Context, start time.Time) LogInfo {
	return LogInfo{
		ClientIP:    c.ClientIP(),
		Date:        start.Format(time.RFC3339),
		Method:      c.Request.Method,
		RequestURI:  c.Request.URL.RequestURI(),
		Referer:     c.Request.Referer(),
		HTTPVersion: c.Request.Proto,
		Size:        c.Writer.Size(),
		Status:      c.Writer.Status(),
		UserAgent:   c.Request.UserAgent(),
		Latency:     time.Now().Sub(start),
	}
}

func ConvertToMapFromBody(c *gin.Context) (m map[string]interface{}, err error) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	if len(b) != 0 {
		err = json.Unmarshal(b, &m)
	}
	return
}
