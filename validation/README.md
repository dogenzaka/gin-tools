Validation
=========

## Installation

```
$ go get github.com/dogenzaka/gin-tools/validation
```

## Requirements

- [gin-gonic](http://gin-gonic.github.io/gin/)

### Tests

- [goconvey](https://github.com/smartystreets/goconvey)

## Example

```go
func main() {

    r := gin.Default()
    
    r.Use(validation.ValidatePathParam("name", validator.MinLength{1}, validator.MaxLength{32}))
    r.GET("/user/:name", func(c *gin.Context) {
        name := c.Params.ByName("name")
        message := "Hello "+name
        c.String(200, message)
    })

    r.Run(":8080")
}
```
