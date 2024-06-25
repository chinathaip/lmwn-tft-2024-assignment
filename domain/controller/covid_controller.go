package controller

import (
	"github.com/chinathaip/lmwn-tft-2024-assignment/domain/repository"
	"github.com/gin-gonic/gin"
)

type CovidController struct {
	covidRepository repository.CovidRepository
}

func NewCovidController(covidRepository repository.CovidRepository) *CovidController {
	return &CovidController{}
}

func (cc *CovidController) HandleSummary(c *gin.Context) {
	cc.covidRepository.Summarize()
}
