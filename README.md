# Middleware for [Forgery][forgery]

## Known issues

## API
### middleware.BodyParser
Currently parses only JSON data.

## Example
```go
package main

import (
  "fmt"
  "github.com/southern/middleware"
  "github.com/ricallinson/forgery"
)

func main() {
  server := f.CreateServer()
  server.Use(middleware.BodyParser)
  server.Get("/", func (req *f.Request, res *f.Response, next func()) {
    if len(req.Request.Map) > 0 {
      fmt.Printf("Map: %+v\n", req.Request.Map)
    }
  })
  server.Listen(8080)
}
```

[forgery]: https://github.com/ricallinson/forgery
