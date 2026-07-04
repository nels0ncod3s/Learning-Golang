package main 

import (
    "fmt"
    
)

func reverse(word string) string {
    reversed := ""

    for i := len(word) - 1; i >= 0; i-- {
        reversed += string(word[i])
    }

    return reversed
}

func isPalindrome(word string) bool {
    reversed := reverse(word)

    return word == reversed
}

func main (){
    var word string

    fmt.Print("Enter a word: ")
    fmt.Scanln(&word)

    if isPalindrome(word) {
        fmt.Println(word, "is a palindrome.")
    } else {
        fmt.Println(word, "is NOT a palindrome.")
    }
}