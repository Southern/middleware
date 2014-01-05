package middleware

import (
  "encoding/json"
  "github.com/ricallinson/stackr"
  "io/ioutil"
)

func BodyParser(req *stackr.Request, res *stackr.Response, next func()) {
  defer next()
  contentType := req.Request.Header.Get("content-type")
  if contentType == "application/json" {
    req.Body = make(map[string]string)
    body, err := ioutil.ReadAll(req.Request.Body)
    if err != nil {
      return
    }
    err = json.Unmarshal(body, &req.Body)
    if err != nil {
      return
    }
  }
}
