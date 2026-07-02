package main 

import (
    "fmt"
)

//Temperature Converter
func main(){
    var temp float64 //Storing user input

    fmt.Print("Enter the temperature in Celcius: ")
    fmt.Scanln(&temp) //Prompting user's input

    Fah := (temp * 1.8) + 32 //conversion logic
    Kel := temp + 273.15 //conversion logic
    
    fmt.Printf("Temperature in Fahrenheit is %.2f\n", Fah)
    fmt.Printf("Temperature in Kelvin is %.2f", Kel)
}