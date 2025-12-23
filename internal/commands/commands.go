package commands

import (
	"expense-tracker/internal/domain"
	"expense-tracker/internal/presenter"
	"expense-tracker/internal/service"
	"fmt"
)

// ------------------
// Icommand
// ------------------
type ICommand interface {
	Validate() error
	Execute() error
}

// ------------------
// Base command
// ------------------
type Command struct {
	service   *service.ExpenseService
	presetner *presenter.ConsolePresenter
}

func newComamnd(s *service.ExpenseService, p *presenter.ConsolePresenter) *Command {
	return &Command{
		service:   s,
		presetner: p,
	}
}

// ------------------
// Add command
// ------------------
type AddCommand struct {
	Command
	description string
	category    string
	amount      float64
}

func NewAddCommand(s *service.ExpenseService, p *presenter.ConsolePresenter, amount float64, desc, categ string) *AddCommand {
	return &AddCommand{
		Command:     *newComamnd(s, p),
		description: desc,
		category:    categ,
		amount:      amount,
	}
}

func (ac *AddCommand) Validate() error {
	if err := domain.ValidateExpense(ac.amount, ac.description, ac.category); err != nil {
		return err
	}

	return nil
}

func (ac *AddCommand) Execute() error {
	if err := ac.service.AddExpense(ac.amount, ac.description, ac.category); err != nil {
		return err
	}

	ac.presetner.Success("Expense added successfully")

	return nil
}

// ------------------
// Update command
// ------------------
type UpdateCommand struct {
	Command
	id          int
	description string
	category    string
	amount      float64
}

func NewUpdateCommand(s *service.ExpenseService, p *presenter.ConsolePresenter, id int, amount float64, desc, categ string) *UpdateCommand {
	return &UpdateCommand{
		Command:     *newComamnd(s, p),
		id:          id,
		description: desc,
		category:    categ,
		amount:      amount,
	}
}

func (uc *UpdateCommand) Validate() error {
	if err := domain.ValidateID(uc.id); err != nil {
		return err
	}

	if uc.amount != NoIDSelected {
		if err := domain.ValidateAmount(uc.amount); err != nil {
			return err
		}
	}

	if uc.description != "" {
		if err := domain.ValidateDescription(uc.description); err != nil {
			return err
		}
	}

	if uc.category != "" {
		if err := domain.ValidateCategory(uc.category); err != nil {
			return err
		}
	}

	return nil
}

func (uc *UpdateCommand) Execute() error {
	if err := uc.service.UpdateExpense(uc.id, uc.amount, uc.description, uc.category); err != nil {
		return err
	}

	uc.presetner.Success(fmt.Sprintf("Expense %d updated successfully", uc.id))

	return nil
}

// -------------------
// Delete command
// -------------------
type DeleteCommand struct {
	Command
	id int
}

func NewDeleteCommand(s *service.ExpenseService, p *presenter.ConsolePresenter, id int) *DeleteCommand {
	return &DeleteCommand{
		Command: *newComamnd(s, p),
		id:      id,
	}
}

func (dc *DeleteCommand) Validate() error {
	if err := domain.ValidateID(dc.id); err != nil {
		return err
	}

	return nil
}

func (dc *DeleteCommand) Execute() error {
	if err := dc.service.DeleteExpense(dc.id); err != nil {
		return err
	}

	dc.presetner.Success(fmt.Sprintf("Expense %d deleted successfully", dc.id))

	return nil
}

// -------------------
// Summary command
// -------------------
type SummaryCommand struct {
	Command
	monthID  int
	category bool
}

func NewSummaryCommand(s *service.ExpenseService, p *presenter.ConsolePresenter, monthID int, category bool) *SummaryCommand {
	return &SummaryCommand{
		Command:  *newComamnd(s, p),
		category: category,
		monthID:  monthID,
	}
}

func (sc *SummaryCommand) Validate() error {
	if sc.monthID != NoIDSelected {
		if err := domain.ValidateMonthID(sc.monthID); err != nil {
			return err
		}
	}

	return nil
}

func (sc *SummaryCommand) Execute() error {
	switch {
	case sc.monthID != NoIDSelected:
		sc.presetner.ShowSummaryofMonth(sc.monthID)
	case sc.category:
		sc.presetner.ShowSummaryofCategory()
	default:
		sc.presetner.ShowTotalSummary()
	}

	return nil
}

// -------------------
// List Command
// -------------------
type ListCommand struct {
	Command
	monthID  int
	category string
}

func NewListCommand(s *service.ExpenseService, p *presenter.ConsolePresenter, monthID int, category string) *ListCommand {
	return &ListCommand{
		Command:  *newComamnd(s, p),
		monthID:  monthID,
		category: category,
	}
}

func (lc *ListCommand) Validate() error {
	if lc.monthID != NoIDSelected {
		if err := domain.ValidateMonthID(lc.monthID); err != nil {
			return err
		}
	}

	if lc.category != "" {
		if err := domain.ValidateCategory(lc.category); err != nil {
			return err
		}
	}

	return nil
}

func (lc *ListCommand) Execute() error {
	switch {
	case lc.monthID != NoIDSelected:
		lc.presetner.ShowList(lc.service.GetExpensesofMonth(lc.monthID))
	case lc.category != "":
		lc.presetner.ShowList(lc.service.GetExpensesByCategory(lc.category))
	default:
		lc.presetner.ShowList(lc.service.GetAllExpenes())
	}

	return nil
}
