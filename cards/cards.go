package cards

import (
	"card-management/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type AllCards map[string][]*Card

func (a *AllCards) String() string {
	return "all cards data"
}

func (a *AllCards) Set(value string) error {
	if *a == nil {
		// init the map for cards
		*a = make(map[string][]*Card)
	}
	card_details := strings.Split(strings.Trim(value, "\n"), "-")
	if len(card_details) < 4 {
		err_str := fmt.Sprintf("input value : %v doesn't match the expected format", value)
		return errors.New(err_str)
	}

	// fmt.Println(card_details)
	issuer := card_details[0]
	category := card_details[1]
	discount_categories := utils.StringToMapInt(card_details[2])
	card_benefits := utils.StringToMapString(card_details[3])
	new_card := NewCard(issuer, category, discount_categories, card_benefits)

	// compare if two structs are same and skip adding the record
	_, ok := (*a)[issuer]
	if ok {
		for _, val := range (*a)[issuer] {
			if reflect.DeepEqual(val, new_card) {
				fmt.Println("skipping.. as data available for")
				val.PrintCardInfo()
				return nil
			}
		}
	}

	(*a)[issuer] = append((*a)[issuer], new_card)
	// fmt.Printf("%v", (*a)[0])
	return nil
}

type Card struct {
	CardIssuer         string            `json:"card_issuer"`
	CardCategory       string            `json:"card_category"`
	DiscountCategories map[string]int    `json:"discount_categories"`
	CardBenefits       map[string]string `json:"card_benefits"`
}

func NewCard(issuer string, category string, discount_categories map[string]int, card_benefits map[string]string) *Card {
	return &Card{
		CardIssuer:         issuer,
		CardCategory:       category,
		DiscountCategories: discount_categories,
		CardBenefits:       card_benefits,
	}
}

func (card *Card) GetCategory() string {
	return card.CardCategory
}

func (card *Card) GetCardIssuer() string {
	return card.CardIssuer
}

func (card *Card) GetDiscount(category string) (int, error) {
	if _, ok := card.DiscountCategories[category]; ok {
		return card.DiscountCategories[category], nil
	} else {
		error := fmt.Sprintf("Category %v not available", category)
		return 0, errors.New(error)
	}
}

func (card *Card) PrintCardInfo() {
	fmt.Printf("%v, %v, %v, %v \n", card.CardIssuer, card.CardCategory, card.DiscountCategories, card.CardBenefits)

}

func (card *Card) GetJson() ([]byte, error) {
	return json.MarshalIndent(card, "", "\t")
}

func UnmarshallData(all_cards *AllCards) {
	// unmarshal the total data that is already saved in prev files
	content, err := os.ReadFile("card-app.json")
	if err != nil {
		fmt.Println(err)
	}
	err_json := json.Unmarshal(content, all_cards)

	if err_json != nil {
		fmt.Println(err_json)
	}
}
