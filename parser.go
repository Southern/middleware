package middleware

import (
  "encoding/json"
  "github.com/ricallinson/stackr"
  "io/ioutil"
)

func BodyParser(req *stackr.Request, res *stackr.Response, next func()) {
  var j map[string]interface{}
  defer next()
  contentType := req.Header.Get("content-type")
  if contentType[0:16] == "application/json" {
    if req.Map == nil {
      req.Map = make(map[string]interface{})
    }
    req.Map["json"] = make(map[string]interface{})
    body, err := ioutil.ReadAll(req.Request.Body)
    if err != nil {
      return
    }
    err = json.Unmarshal(body, &j)
    if err != nil {
      return
    }
    req.Map["json"] = j
  }
}
