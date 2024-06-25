package controller

import (
	"net/http"

	"github.com/chinathaip/lmwn-tft-2024-assignment/domain/repository"
	"github.com/gin-gonic/gin"
)

type CovidController struct {
	covidRepository repository.CovidRepository
}

func NewCovidController(covidRepository repository.CovidRepository) *CovidController {
	return &CovidController{covidRepository: covidRepository}
}

func (cc *CovidController) HandleSummary(c *gin.Context) {
	c.JSON(http.StatusOK, cc.covidRepository.Summarize())
}
