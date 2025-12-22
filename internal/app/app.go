package app

import (
	"expense-tracker/internal/commands"
	"expense-tracker/internal/presenter"
	"expense-tracker/internal/service"
)

type App struct {
	service   *service.ExpenseService
	presenter *presenter.ConsolePresenter
}

func NewApp(s *service.ExpenseService, p *presenter.ConsolePresenter) *App {
	return &App{service: s, presenter: p}
}

func (a *App) createCommand(commandName string, flag *commands.CmdFlags) commands.ICommand {
	switch commandName {
	case "add":
		return commands.NewAddCommand(a.service, a.presenter, flag.Amount, flag.Description, flag.Category)
	case "delete":
		return commands.NewDeleteCommand(a.service, a.presenter, flag.Delete)
	case "update":
		return commands.NewUpdateCommand(a.service, a.presenter, flag.Update, flag.Amount, flag.Description, flag.Category)
	case "list":
		return commands.NewListCommand(a.service, a.presenter, flag.MonthID, flag.Category)
	case "summary":
		return commands.NewSummaryCommand(a.service, a.presenter, flag.MonthID, flag.Cate)
	default:
		return nil
	}
}

// Run the application
func (a *App) Run(args []string) error {
	// parse flags
	flags, err := commands.ParseFlags(args)
	if err != nil {
		return err
	}

	// determine command
	commandName, err := flags.DetermineCommand()
	if err != nil {
		return err
	}

	// create command
	command := a.createCommand(commandName, flags)

	// validate
	if err := command.Validate(); err != nil {
		return err
	}

	return command.Execute()
}
