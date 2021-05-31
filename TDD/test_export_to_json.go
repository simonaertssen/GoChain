package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
)


type User struct {
    name     string
    digest   string
    balance  float32
}


func main() {
    U1 := User {
                name: "Sean",
                digest: "40d18d5a7ae85f9597a40f1306041fd1", 
                balance: 50.23,
                }
    U2 := User {
                name: "Donnie",
                digest: "6f171d413bee711762beff4276595068", 
                balance: 14572.291,
                }
    var all_users = []User {U1, U2}
    fmt.Println(all_users)

    data, error := json.Marshal(all_users)
    if error != nil {
        fmt.Println("error:", error)
    }

    error = ioutil.WriteFile("test_users.json", data, 0644)
}
