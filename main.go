package main

import (
	"card-management/cards"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	card := cards.NewCard("chase", "flex", map[string][]int{"travel": {2, 3, 4}})

	discount, err := card.GetDiscount("travel")
	if err != nil {
		fmt.Println("error received")
		fmt.Println(err)
	} else {
		fmt.Println(discount)
	}
}
