gin-tools
=========

[![Build Status](http://img.shields.io/travis/dogenzaka/gin-tools.svg?style=flat)](https://travis-ci.org/dogenzaka/gin-tools)
[![Coverage](http://img.shields.io/codecov/c/github/dogenzaka/gin-tools.svg?style=flat)](https://codecov.io/github/dogenzaka/gin-tools)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/dogenzaka/gin-tools/blob/master/LICENSE)

gin-tools is middlewares for [gin-gonic](http://gin-gonic.github.io/gin/)

## Installation

- Validation
```
$ go get github.com/dogenzaka/gin-tools/validation
```

## Requirements

- [gin-gonic](http://gin-gonic.github.io/gin/)

### Tests

- [goconvey](https://github.com/smartystreets/goconvey)

### Example

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

## License
gin-tools is licensed under the MIT.
