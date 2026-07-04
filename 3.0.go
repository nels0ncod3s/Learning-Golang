package main 

import (
    "fmt"
)

//reverse a word
func reverse(word string) string {
    reversed := ""

    for i := len(word) - 1; i >= 0; i-- {
        reversed += string(word[i])
    }

    return reversed
}

func main (){
    var word string

    fmt.Print("Enter a word: ")
    fmt.Scanln(&word)

    result := reverse(word)

    fmt.Println("Reversed word:", result)
}