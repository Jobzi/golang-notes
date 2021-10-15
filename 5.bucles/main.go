package main

import (
	"fmt"
)

//La palabra clave se utiliza para iterar más fácilmente sobre una matriz, división o mapa. Devuelve tanto el índice como el valor.range
func forWithRange() {
  fmt.Println("FOR WITH RANGE")
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }

  forWithRange()
}