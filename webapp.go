package main

import (
	"card-management/cards"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func getAllcards(rw http.ResponseWriter, r *http.Request) {
	var allCards = make(cards.AllCards)
	cards.UnmarshallData(&allCards)
	path := r.URL.Path
	fmt.Println(strings.Trim(path, "/"))
	// fmt.Println(path)
	switch path {
	case "/":
		fmt.Fprintf(rw, "Welcome to webapp")
	case "/all":
		var all_cards []cards.Card
		for _, cards := range allCards {
			for _, card := range cards {
				all_cards = append(all_cards, *card)

			}
		}
		// fmt.Println(all_cards)
		var byt []byte
		// var err error
		byt, _ = json.MarshalIndent(all_cards, "", "\t")
		fmt.Fprintf(rw, string(byt))
	case "/travel":
		var all_cards []*cards.Card
		for _, cards := range allCards {
			all_cards = append(all_cards, cards...)
		}
		sort.Slice(all_cards, func(i, j int) bool {
			return all_cards[i].DiscountCategories["travel"] > all_cards[j].DiscountCategories["travel"]
		})
		for _, card := range all_cards {
			fmt.Println(*card)
		}
	default:
		var byt []byte
		var err error
		if allCards[path[1:]] != nil {
			byt, err = json.MarshalIndent(allCards[path[1:]], "", "\t")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			byt = []byte("[]")
		}
		fmt.Fprint(rw, string(byt))
	}

}

func main() {
	fmt.Println("welcome to webapp")
	optPort := flag.Int("port", 8080, "Provide the port number that has to used")
	flag.Parse()
	fmt.Println(*optPort)

	// declare port
	port := ":" + strconv.Itoa(*optPort)
	// define web routes
	http.HandleFunc("/", getAllcards)
	http.HandleFunc("/abc", func(rw http.ResponseWriter, r *http.Request) { fmt.Fprintf(rw, "Welcome to %s", r.URL.Path) })

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Print(err)
		return
	}
}
