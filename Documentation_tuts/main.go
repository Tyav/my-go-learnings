package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println("We are testing this thing sha")

	// we want to import a module
	fmt.Println(quote.Go())
	
}