package main 

import (
    "fmt"
    "os"
)

//File reader using defer lol
func main() {
    file, err := os.Open("note.txt")

    if err != nil{
        fmt.Println(err)
        return
    }

    defer file.Close()

    fmt.Println("File opened successfully")
}