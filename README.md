# Middleware for [Forgery][forgery]

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
    // JSON object with string keys and string values
    json := req.Map["json"].(map[string]string)
    if len(json) > 0 {
      fmt.Printf("JSON: %+v\n", json)
    }
  })
  server.Listen(8080)
}
```

[forgery]: https://github.com/ricallinson/forgery
