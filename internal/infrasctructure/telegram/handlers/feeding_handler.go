package handlers

import (
	"github.com/qrave1/PetFeedingBot/internal/usecase"
	tele "gopkg.in/telebot.v4"
)

type FeedingHandler interface {
	Add() tele.HandlerFunc
	GetFeedings() tele.HandlerFunc
	Delete() tele.HandlerFunc
}

type FeedingHandlerImpl struct {
	feedingUsecase usecase.FeedingUsecase
}

func NewFeedingHandlerImpl(feedingUsecase usecase.FeedingUsecase) *FeedingHandlerImpl {
	return &FeedingHandlerImpl{feedingUsecase: feedingUsecase}
}

//func (f *FeedingHandlerImpl) Add() tele.HandlerFunc {
//	return func(c tele.Context) error {
//
//	}
//}
//
//func (f *FeedingHandlerImpl) GetFeedings() tele.HandlerFunc {
//	return func(c tele.Context) error {
//
//	}
//}
//
//func (f *FeedingHandlerImpl) Delete() tele.HandlerFunc {
//	return func(c tele.Context) error {
//
//	}
//}
