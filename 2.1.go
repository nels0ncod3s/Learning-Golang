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
    fmt.Println("1 - Easy, 2 - Medium, 3 - Hard")
    fmt.Scanln(&choice)

    //conditional logic
    if choice == 1{ //Logic for Easy
        attempt := 5
    
        fmt.Print("")
        fmt.Println("You have chosen Easy!")
        fmt.Println("1 - 10")
        fmt.Printf("You have %d attempts\n", attempt)
        fmt.Println("")

        randnum := rand.Intn(10) + 1 //generating a random number between 1 and 10
        

        for attempt > 0 {

            var userChoice int

            fmt.Print("Guess the number:")
            fmt.Scanln(&userChoice)
            

            if userChoice == randnum {
                fmt.Printf("You guessed the magic number, it was %d", randnum)
                break //To stop the loop when you win 
            } else {
                attempt--

                if attempt > 1{
                    fmt.Printf("%d attemept(s) left\n", attempt)
                    fmt.Println("Keep Trying!")
                    fmt.Println("")
                } else if attempt == 1{
                    fmt.Println("LAST CHANCE!!!")
                    fmt.Println("")
                } else{
                    fmt.Println("Game Over!")
                    break
                }
            }
        }
    } else if choice == 2{ // Logic for Medium
        
        attempt := 7
    
        fmt.Print("")
        fmt.Println("You have chosen Medium!")
        fmt.Println("1 - 50")
        fmt.Printf("You have %d attempts\n", attempt)
        fmt.Println("")

        randnum := rand.Intn(50) + 1 //generating a random number between 1 and 50
        

        for attempt > 0 {

            var userChoice int

            fmt.Print("Guess the number:")
            fmt.Scanln(&userChoice)
            

            if userChoice == randnum {
                fmt.Printf("You guessed the magic number, it was %d", randnum)
                break //To stop the loop when you win 
            } else {
                attempt--

                if attempt > 1{
                    fmt.Printf("%d attemept(s) left\n", attempt)
                    fmt.Println("Keep Trying!")
                    fmt.Println("")
                } else if attempt == 1{
                    fmt.Println("LAST CHANCE!!!")
                    fmt.Println("")
                } else{
                    fmt.Println("Game Over!")
                    break
                }
            }
        }
    } else if choice == 3{ //Logic for Hard
        attempt := 10
    
        fmt.Print("")
        fmt.Println("You have chosen Hard!")
        fmt.Println("1 - 100")
        fmt.Printf("You have %d attempts\n", attempt)
        fmt.Println("")

        randnum := rand.Intn(100) + 1 //generating random number between 1 and 100
        

        for attempt > 0 {

            var userChoice int

            fmt.Print("Guess the number:")
            fmt.Scanln(&userChoice)
            

            if userChoice == randnum {
                fmt.Printf("You guessed the magic number, it was %d", randnum)
                break //To stop the loop when you win 
            } else {
                attempt--

                if attempt > 1{
                    fmt.Printf("%d attemept(s) left\n", attempt)
                    fmt.Println("Keep Trying!")
                    fmt.Println("")
                } else if attempt == 1{
                    fmt.Println("LAST CHANCE!!!")
                    fmt.Println("")
                } else{
                    fmt.Println("Game Over!")
                    break
                }
            }
        }
    }
    
}