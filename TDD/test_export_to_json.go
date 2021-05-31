package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
)


type User struct {
    Name     string
    Digest   string
    Balance  float32
}


func main() {
    U1 := User {
                Name:   "Sean",
                Digest: "40d18d5a7ae85f9597a40f1306041fd1", 
                Balance: 50.23,
                }
    U2 := User {
                Name:   "Donnie",
                Digest: "6f171d413bee711762beff4276595068", 
                Balance: 14572.291,
                }

    var all_users = []User {U1, U2}
    //fmt.Println(all_users)

    file, error := json.MarshalIndent(all_users, "", " ")
    if error != nil {
        fmt.Println("error:", error)
    }

    error = ioutil.WriteFile("test_users.json", file, 0644)
}
