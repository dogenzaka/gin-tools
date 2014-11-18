Logging
=========

## Installation

```
$ go get github.com/dogenzaka/gin-tools/logging
```

## Requirements

- [gin-gonic](http://gin-gonic.github.io/gin/)

### Tests

- [goconvey](https://github.com/smartystreets/goconvey)

## Example

```go
func main() {

    r := gin.Default()
    
    r.Use(logging.AccessLogger(os.Stdout))
    r.GET("/user/:name", func(c *gin.Context) {
        name := c.Params.ByName("name")
        message := "Hello "+name
        c.String(200, message)
    })

    r.Run(":8080")
}
```
