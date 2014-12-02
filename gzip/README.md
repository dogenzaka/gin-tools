gzip middleware for gin
====

This middleware comes from original gin-gonic middleware located in https://github.com/gin-gonic/contrib.
Since gin does not reset ResponseWriter in context when overwritten, this middleware resolves issue
by recovering original writer to the context.

```go
package main

import (
	"fmt"
	"github.com/dogenzaka/gin-tools/gzip"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```
