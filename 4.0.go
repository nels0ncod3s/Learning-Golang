package main

import (
 "fmt"
 "os"
 "strings"
)

func main() {

 data, err := os.ReadFile("story.txt")

 if err != nil {
  fmt.Println("Could not read file.")
  return
 }

 text := string(data)

 words := strings.Fields(text)

 counts := make(map[string]int)

 for _, word := range words {
  counts[word]++
 }

 fmt.Println("------ WORD COUNT ------")

 for word, count := range counts {
  fmt.Println(word, ":", count)
 }
}