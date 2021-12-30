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
	flag.Var(&allCards, "card-info", "Card information with '-' seperated values")

	// parse the options from command line
	flag.Parse()

	fmt.Println("Thanks for using, Card Management App", *optDisplay, *optSaveFile)

	// save the card information to a file
	if len(*optSaveFile) > 0 {
		cwd, _ := os.Getwd()
		*optSaveFile = filepath.Join(cwd, *optSaveFile)
		fmt.Println("Saving file to ", *optSaveFile, "location")

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

	// unmarshal the total data that is already saved in prev files
	var unmarshal_data cards.AllCards
	content, err := os.ReadFile("card-app.json")
	if err != nil {
		fmt.Println(err)
	}
	// buf_reader := bufio.NewReader(read_file)
	// var read_bytes = []byte{}
	// bytes, err := buf_reader.Read(read_bytes)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("total bytes read:", bytes, read_bytes)
	err_json := json.Unmarshal(content, &unmarshal_data)
	if err_json != nil {
		fmt.Println(err)
	}
	allCards = append(allCards, unmarshal_data...)

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
}
