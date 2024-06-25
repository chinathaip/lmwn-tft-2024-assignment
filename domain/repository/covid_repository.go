package repository

import "github.com/chinathaip/lmwn-tft-2024-assignment/domain/dto"

type CovidRepository interface {
	Summarize() dto.CovidSummary
}
