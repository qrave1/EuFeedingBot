package main

import (
	"log"
	"os"
	"time"

	"github.com/qrave1/PetFeedingBot/internal/bot"
	"github.com/qrave1/PetFeedingBot/internal/repository"
	"github.com/qrave1/PetFeedingBot/internal/usecase"

	"github.com/jmoiron/sqlx"
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

	db, err := sqlx.Connect("sqlite3", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return
	}

	petRepo := repository.NewPetRepo(db)

	petUsecase := usecase.NewPetUsecaseImpl(petRepo)

	_ = bot.NewEuFeedingBot(b, petUsecase)

	b.Start()
}
