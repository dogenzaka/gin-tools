// this gzip middleware comes from github.com/gin-gonic/contrib/gzip
// Fixed response writer issue
package gzip

import (
	gogzip "compress/gzip"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	BestCompression    = gogzip.BestCompression
	BestSpeed          = gogzip.BestSpeed
	DefaultCompression = gogzip.DefaultCompression
	NoCompression      = gogzip.NoCompression

	headerAcceptEncoding  = "Accept-Encoding"
	headerContentEncoding = "Content-Encoding"
	headerContentLength   = "Content-Length"
	headerVary            = "Vary"

	encodingGzip = "gzip"
)

type gzipWriter struct {
	gin.ResponseWriter
	gzwriter *gogzip.Writer
}

func newGzipWriter(writer gin.ResponseWriter, gzWriter *gogzip.Writer) *gzipWriter {
	return &gzipWriter{writer, gzWriter}
}

func (g *gzipWriter) Write(data []byte) (int, error) {
	return g.gzwriter.Write(data)
}

func Gzip(level int) gin.HandlerFunc {

	return func(c *gin.Context) {

		req := c.Request
		// Skip if accept-encoding header does not contain gzip
		if !strings.Contains(req.Header.Get(headerAcceptEncoding), encodingGzip) {
			c.Next()
			return
		}

		writer := c.Writer
		defer func() {
			// Recover original response writer
			c.Writer = writer
		}()

		// Initiate gzip writer
		gz, err := gogzip.NewWriterLevel(writer, level)
		if err != nil {
			c.Fail(500, err)
			return
		}
		defer gz.Close()

		headers := writer.Header()
		headers.Set(headerContentEncoding, encodingGzip)
		headers.Set(headerVary, headerAcceptEncoding)

		c.Writer = newGzipWriter(writer, gz)
		c.Next()
		headers.Del(headerContentLength)
	}

}
