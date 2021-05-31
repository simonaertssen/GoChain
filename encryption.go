package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "crypto/md5"
)

func main() {
    // Read the state file
    data, error := ioutil.ReadFile("data/state.json")
    if error != nil {
      fmt.Print(err)
    }

    // If users are not encrypted yet, encrypt
    s := "Simon"
    digest := md5.Sum([]byte(s))

    fmt.Printf("%x\n", digest)
}
