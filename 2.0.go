package main 

import (
    "fmt"
)

//FizzBuzz
func main() {
    var num int //storing user's input

    fmt.Print("Enter a number: ")
    fmt.Scanln(&num) //prompting user's input

    //Looping logic
    for i := 1; i < num; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i % 3 == 0 {
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            fmt.Println("Buzz")
        } else{
            fmt.Println(i)
        }
    }
}