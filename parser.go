package middleware

import (
  "encoding/json"
  "github.com/ricallinson/stackr"
  "io/ioutil"
)

func BodyParser(req *stackr.Request, res *stackr.Response, next func()) {
  defer next()
  contentType := req.Header.Get("content-type")
  if contentType[0:16] == "application/json" {
    req.Map = make(map[string]interface{})
    body, err := ioutil.ReadAll(req.Request.Body)
    if err != nil {
      return
    }
    err = json.Unmarshal(body, &req.Map)
    if err != nil {
      return
    }
  }
}
