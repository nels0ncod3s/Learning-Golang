package main 

import (
    "fmt"
)

func main(){

    //creating a map 
	phonebook := map[int]string {
        001 : "Davido",
        121 : "FireService",
        222 : "Nelson",
    }

    fmt.Println(phonebook)

    phonebook[248]= "Sam"
    
    fmt.Println(phonebook)
}