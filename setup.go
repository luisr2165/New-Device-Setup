package main

import (
  "fmt"
)

func main() {

  var name string

  fmt.Println("What is your name?")

  fmt.Scan(&name)

  fmt.Println("Hello", name)

  var age int

  fmt.Println("What is your age?")

  fmt.Scan(&age)

  fmt.Printf("The name of the user is %s and they are %d years old.", name, age)

  newMessage := fmt.Sprintf("Their name is %s and their age is %d years old", name, age)

  fmt.Println(newMessage)
  
}
