package main

import (
    "bytes"
    json2 "encoding/json"
    "fmt"
    "github.com/joho/godotenv"
    "os"
)
import "net/http"

type Message struct {
    Usename string `json:"username"`
    Content string `json:"content"`
}


func main() {
    err := godotenv.Load(".env")
    
    if err != nil {
        panic(err)
    }

    client := http.Client{}

    json, err := json2.Marshal(Message { Usename: "do it, cunt", Content: "make github push"})

    if err != nil {
        panic(err)
    }

    fmt.Println(string(json))


    fmt.Println(os.Getenv("token"))
    url := os.Getenv("token")

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

    req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/json")

    if err != nil {
        panic(err)
    }

    resp, err := client.Do(req)

    if err != nil {
        return
    }

    fmt.Println(resp)
}