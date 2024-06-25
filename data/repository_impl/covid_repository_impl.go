package repository_impl

import "github.com/chinathaip/lmwn-tft-2024-assignment/domain/dto"

type CovidRepositoryImpl struct {
}

func NewCovidRepositoryImpl() *CovidRepositoryImpl {
	return &CovidRepositoryImpl{}
}

func (cr *CovidRepositoryImpl) Summarize() *dto.Covid {
	return nil
}
