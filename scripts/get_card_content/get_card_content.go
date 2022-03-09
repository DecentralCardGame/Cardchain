package main

import (
  "encoding/json"
  "os"
  "fmt"
)

type Content struct {
  Content []byte
}

func main() {
  contentString := fmt.Sprintf(`{"Content": "%s"}`, os.Args[1])

  var content Content
  err := json.Unmarshal([]byte(contentString), &content)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(content.Content))
}
