# Middleware for [Forgery][forgery]

## Known issues
- BodyParser only parses strings at this time.

## API
### middleware.BodyParser
Currently parses only JSON data.

## Example
```go
package main

import (
  "github.com/southern/middleware"
  "github.com/ricallinson/forgery"
)

func main() {
  server := f.CreateServer()
  server.Use(middleware.BodyParser)
  server.Listen(8080)
}
```

[forgery]: https://github.com/ricallinson/forgery
