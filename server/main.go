package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte("<h1>API Interface</h1>"))
}


func main() {
  walk()
  fs := http.FileServer(http.Dir("../client/dist"))
  http.Handle("/", fs)
  http.HandleFunc("/api", ApiHandler)

  log.Println("Starting HTTP server on port 8080")

  err := http.ListenAndServe(":8080", nil)

  if err != nil {
    log.Fatal(err)
  }
}

func walk() {
  walkFn := func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    fmt.Println(path)
    return nil
  }

  err := filepath.Walk("../", walkFn)
  if err != nil {
    fmt.Println(err)
  }
}
