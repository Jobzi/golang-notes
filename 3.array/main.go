package main

import "fmt"

//Los slices son muy parecido al array, sin embargo son mas flexibles y faciles de usar
func sliceExample() {
  fmt.Println("EJEMPLO DE COMO USUR LOS SLICES")
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}

func main() {
	numberList:= [...]int{0,1,2,3,4,5,6}
	namesList:=[...]string{"a","b","c","d","e","f","g","h","i","j"}

	arr1 := [5]int{} //not initialized
  	arr2 := [5]int{1,2} //partially initialized
  	arr3 := [5]int{1,2,3,4,5} //fully initialized
	
	fmt.Println(numberList)
	fmt.Println(namesList)

	fmt.Println(arr1)// no inizializado output:[0 0 0 0 0]
	fmt.Println(arr2)// parcialemte Inizialado output:[1 2 0 0 0]
	fmt.Println(arr3)// inizializado output:[1 2 3 4 5]

	sliceExample()
}