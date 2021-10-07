package main

import (
	"fmt"
	"os"
)

func main() {
	//ErrorF like used
	const lastname, id = "bueller", 17
	err1 := fmt.Errorf("user %q (id %d) not found", lastname, id)
	fmt.Println(err1.Error())


	//Fprintf like used
	const name, age = "Kim", 22
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)
	// The n and err return values from Fprintf are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)
}