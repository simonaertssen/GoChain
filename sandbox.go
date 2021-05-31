package main

import (
    "fmt"
    "crypto/md5"
)

func main() {
    s := "Simon"
    digest := md5.Sum([]byte(s))

    fmt.Printf("%x\n", digest)
}
