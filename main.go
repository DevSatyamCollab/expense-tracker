package main

import (
	"expense-tracker/internal/app"
	"expense-tracker/internal/domain"
	"expense-tracker/internal/presenter"
	"expense-tracker/internal/service"
	"expense-tracker/internal/storage"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// initialize storage
	jsonStorage, err := storage.GetStorage()
	if err != nil {
		return err
	}

	// initialize tracker
	tracker := domain.GetExpenseTracker()

	// load the existing file
	if err := jsonStorage.Load(tracker); err != nil {
		return err
	}

	// initialize the service layer
	service := service.NewExpenseService(tracker, jsonStorage)

	// initialize the presenter
	presenter := presenter.NewConsolePresenter(service)

	// run the app
	app := app.NewApp(service, presenter)
	if err := app.Run(os.Args[1:]); err != nil {
		return err
	}

	return nil
}
