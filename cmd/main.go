package main

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/qrave1/PetFeedingBot/cmd/application/config"
	"github.com/qrave1/PetFeedingBot/internal/infrasctructure/telegram"
	"github.com/qrave1/PetFeedingBot/internal/infrasctructure/telegram/handlers"
	"github.com/qrave1/PetFeedingBot/internal/infrasctructure/telegram/presenter"
	"github.com/qrave1/PetFeedingBot/internal/repository"
	"github.com/qrave1/PetFeedingBot/internal/usecase"

	"github.com/jmoiron/sqlx"
	tele "gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)

		return
	}

	sett := tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(sett)
	if err != nil {
		log.Fatal(err)

		return
	}

	b.Use(middleware.AutoRespond())

	db, err := sqlx.ConnectContext(ctx, "sqlite", cfg.DBPath)
	if err != nil {
		log.Fatal(err)

		return
	}

	petRepo := repository.NewPetRepo(db)

	petUsecase := usecase.NewPetUsecaseImpl(petRepo)

	petPresenter := presenter.NewPetPresenter()

	petHandler := handlers.NewPetHandlerImpl(petUsecase, petPresenter)

	_ = telegram.NewPetFeedingBot(b, petHandler)

	// TODO: перенести в app.Start()
	// старт поллинга бота
	go b.Start()

	select {
	case <-ctx.Done():
		slog.Info("Shutting down application...")

		// TODO: тут должен быть просто app.Stop()
		b.Stop()
	}
}
