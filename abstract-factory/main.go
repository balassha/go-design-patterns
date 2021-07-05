package main

import "fmt"

func main() {
	adidasFactory := GetSportsFactory("adidas")
	nikeFactory := GetSportsFactory("nike")

	adShoe := adidasFactory.GetShoe()
	adShorts := adidasFactory.GetShorts()

	nkShoe := nikeFactory.GetShoe()
	nkShorts := nikeFactory.GetShorts()

	PrintShoeDetails(adShoe)
	PrintShoeDetails(nkShoe)

	PrintShortsDetails(adShorts)
	PrintShortsDetails(nkShorts)
}

func PrintShoeDetails(shoe iShoe) {
	fmt.Printf("Shoe Name : %s\n", shoe.GetName())
	fmt.Printf("Shoe Size : %d\n", shoe.GetSize())
}

func PrintShortsDetails(shorts iShorts) {
	fmt.Printf("Shorts Name : %s\n", shorts.GetName())
	fmt.Printf("Shorts Size : %s\n", shorts.GetSize())
}
