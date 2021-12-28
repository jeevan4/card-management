package cards

import (
	"errors"
	"fmt"
)

type Card struct {
	CardIssuer         string
	CardCategory       string
	DiscountCategories map[string][]int
}

func NewCard(issuer string, category string, discount_categories map[string][]int) *Card {
	return &Card{
		CardIssuer:         issuer,
		CardCategory:       category,
		DiscountCategories: discount_categories,
	}
}

func (card *Card) GetCategory() string {
	return card.CardCategory
}

func (card *Card) GetCardIssuer() string {
	return card.CardIssuer
}

func (card *Card) GetDiscount(category string) ([]int, error) {
	if _, ok := card.DiscountCategories[category]; ok {
		return card.DiscountCategories[category], nil
	} else {
		error := fmt.Sprintf("Category %v not available", category)
		return nil, errors.New(error)
	}
}
