package main 

import (
    "fmt"
    //"math/rand"
)

//Number Guessing Game
func main() {
    var choice int //storing user's input
    //guessNum := rand.Intn(10) + 1

    fmt.Println("### Welcome to the Number Guessing Game ###")
    fmt.Print("")
    fmt.Println("Choose Your Difficulty")
    fmt.Print("1 - Easy, 2 - Medium, 3 - Hard")
    fmt.Scanln(&choice)
    fmt.Printf("You have chosen %d", choice)

    //conditional logic
    if choice == 1{
        fmt.Println("You have chosen Easy!")
    } else if choice == 2{
        fmt.Println("You have chosen Medium!")
    } else if choice == 3{
        fmt.Println("You have chosen Hard!")
    } else{
        fmt.Println("You have made an invalid selection!")
    }
}