package main

import (
  "net/http"
  "fmt"
)

func main()  {
  http.HandleFunc("/", homePage)
  http.ListenAndServe(":5000", nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
  if req.URL.Path != "/"{
    http.NotFound(res, req)
    return
  }
  fmt.Fprintf(res, "your server is up")
}
