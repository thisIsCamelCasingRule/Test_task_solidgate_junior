package service

import (
	"Test_task_solidgate_junior/internal/models"
	"regexp"
	"strconv"
	"time"
)

type ServCard interface {
	ValidateCard(card models.CardModel) error
}

func (s Service) ValidateCard(card models.CardModel) error {
	year, err := strconv.Atoi(card.ExpirationYear)
	if err != nil {
		return models.NewValidateCardError("Invalid year", models.CodeValidateCardInvalidYear)
	}

	month, err := strconv.Atoi(card.ExpirationMonth)
	if err != nil || month < 1 || 12 < month {
		return models.NewValidateCardError("Invalid month", models.CodeValidateCardInvalidMonth)
	}

	if year < time.Now().UTC().Year() {
		return models.NewValidateCardError("Card has expired", models.CodeValidateCardExpiredCard)
	}

	if year == time.Now().UTC().Year() && month < int(time.Now().UTC().Month()) {
		return models.NewValidateCardError("Card has expired", models.CodeValidateCardExpiredCard)
	}

	if len(card.CardNumber) != 16 {
		return models.NewValidateCardError("Invalid card number", models.CodeValidateCardInvalidNumber)
	}

	re := regexp.MustCompile(`^4[0-9]+$`)

	if !re.MatchString(card.CardNumber) {
		return models.NewValidateCardError("Invalid card number", models.CodeValidateCardInvalidNumber)
	}

	return nil
}
