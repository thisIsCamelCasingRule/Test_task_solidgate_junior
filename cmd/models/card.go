package models

import "Test_task_solidgate_junior/internal/models"

type CardRequest struct {
	CardNumber      string `json:"card_number" form:"card_number"`
	ExpirationMonth string `json:"expiration_month" form:"expiration_month"`
	ExpirationYear  string `json:"expiration_year" form:"expiration_year"`
}

func (c *CardRequest) ToCardModel() models.CardModel {
	return models.CardModel{
		CardNumber:      c.CardNumber,
		ExpirationMonth: c.ExpirationMonth,
		ExpirationYear:  c.ExpirationYear,
	}
}

type CardValidateSuccessResponse struct {
	Valid bool
}

type CardValidateErrorResponse struct {
	Valid bool
	Error error
}
