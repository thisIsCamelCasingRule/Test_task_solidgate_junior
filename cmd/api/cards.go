package api

import (
	"Test_task_solidgate_junior/cmd/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CardHandler interface {
	ValidateCard(c *gin.Context)
}

func (a Api) ValidateCard(c *gin.Context) {
	var card models.CardRequest
	err := c.ShouldBind(&card)
	if err != nil {
		response := models.CardValidateErrorResponse{
			Valid: false,
			Error: err,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	cardModel := card.ToCardModel()

	err = a.srv.ValidateCard(cardModel)
	if err != nil {
		response := models.CardValidateErrorResponse{
			Valid: false,
			Error: err,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, models.CardValidateSuccessResponse{Valid: true})
}
