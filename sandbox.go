package main

import (
    "time"
    "fmt"
)

func main() {
    a := makeTimestamp()

    fmt.Printf("%d \n", a)
}

func makeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}
