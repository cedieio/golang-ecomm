package main

import "fmt"

func main() {
	fmt.Println("Creating product : ")

	prod := product{Name: "T-Shirt", PriceOne: 32.69}
	fmt.Println("Displaying product")
	fmt.Println(prod.Name)
	fmt.Println(prod.PriceOne)

}

type product struct {
	Name     string
	PriceOne float32
}
