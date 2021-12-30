package main

import (
	"bufio"
	"card-management/cards"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// declare the options
	optDisplay := flag.Bool("display", false, "Display the credit card data that is already stored")
	optSaveFile := flag.String("save-to", "", "File to save the credit card data")
	var allCards cards.AllCards

	// load initial data into allCards before option validation
	cards.UnmarshallData(&allCards)

	flag.Var(&allCards, "card-info", "Card information with '-' seperated values")

	// parse the options from command line
	flag.Parse()

	// save the card information to a file
	if len(*optSaveFile) > 0 {
		cwd, _ := os.Getwd()
		*optSaveFile = filepath.Join(cwd, *optSaveFile)
		fmt.Println("Saving data to ", *optSaveFile, "location")

		byt, err := json.MarshalIndent(allCards, "", "\t")
		fd, err := os.Create(*optSaveFile)

		defer fd.Close()
		buf_writer := bufio.NewWriter(fd)
		buf_writer.Write(byt)
		defer buf_writer.Flush()

		if err != nil {
			fmt.Println(err)
		} else {
			// fmt.Println(string(byt))
		}
	}

	// by default save all data into the file
	byt, err := json.MarshalIndent(allCards, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	fd, err := os.Create("card-app.json")
	if err != nil {
		fmt.Println(err)
	}

	defer fd.Close()
	buf_writer := bufio.NewWriter(fd)
	buf_writer.Write(byt)
	defer buf_writer.Flush()

	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println(string(byt))
	}

	// display all cards information if -display flag is set to true
	if *optDisplay {
		for _, card := range allCards {
			str_byte, _ := card.GetJson()
			fmt.Println(string(str_byte))
		}
	}
	fmt.Println("Thanks for using, Card Management App")
}
