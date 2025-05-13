package main

import (
	"log"
	"os"
	"time"

	"EuFeeding/internal/bot"
	"EuFeeding/internal/repository"
	"EuFeeding/internal/usecase"

	tele "gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

func main() {
	sett := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(sett)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Use(middleware.AutoRespond())

	petRepo := repository.NewPetRepo()

	petUsecase := usecase.NewPetUsecaseImpl()

	_ = bot.NewEuFeedingBot(b, petRepo)

	b.Start()
}
