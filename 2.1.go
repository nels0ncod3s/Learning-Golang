package main 

import (
    "fmt"
    "math/rand"
)

//Number Guessing Game
func main() {
    var choice int //storing user's input

    fmt.Println("### Welcome to the Number Guessing Game ###")
    fmt.Print("")
    fmt.Println("Choose Your Difficulty")
    fmt.Print("1 - Easy, 2 - Medium, 3 - Hard")
    fmt.Scanln(&choice)
    fmt.Printf("You have chosen %d", choice)

    //conditional logic
    if choice == 1{
        fmt.Println("You have chosen Easy!")
        attempt := 5
        radnum := rand.Intn(10) + 1
        var userChoice int
        for i := 0; i < attempt; i++{
            if userChoice == randnum {
                fmt.Println("You guessed the magic number, it was %d", randnum)
                break //To stop the loop when you win 
            } else {
                fmt.Println("Keep Trying!")
            }
        }
    } 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}