package main

import (
  "sha256"
  "os"
  "encoding/hex"
  "fmt"
  "log"
)

func main(){
  bString := os.Args[1]
  b, err := hex.DecodeString(bString)
  if err != nil {
    fmt.Printf("Error with given input: ")
    log.Fatal(err)
  }
  hash := hex.EncodeToString(sha256.CalcDigest(b))
  fmt.Println(hash)
}
