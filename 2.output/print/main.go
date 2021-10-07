package main

import (
	"fmt"
)

func main() {
  var i string = "Hello"
  var j int = 15
  
  //print in only line
  fmt.Print(i,j,"\n") 
 
  //print with format check readme for diferents formant
  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T\n", j, j)

  //print with salt line
  fmt.Println(i,j)
  

}