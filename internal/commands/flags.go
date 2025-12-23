package commands

import (
	"flag"
	"fmt"
	"os"
)

const NoIDSelected = -1

type CmdFlags struct {
	Add         bool
	List        bool
	Summary     bool
	Cate        bool
	Delete      int
	Update      int
	MonthID     int
	Description string
	Category    string
	Amount      float64
}

func ParseFlags(args []string) (*CmdFlags, error) {
	flags := &CmdFlags{}
	fs := flag.NewFlagSet("expense-tracker", flag.ContinueOnError)

	// set custom usage
	fs.Usage = func() {
		usage := `Usage: expense-tracker [options]
Examples:
  expense-tracker -add --desc "Lunch" --amount 20 --categ "transport"
  expense-tracker -upd 0 (--desc "Breakfast" / --amount 25 / --categ "Food"
  expense-tracker -del 0
  expense-tracker -list
  expense-tracker -list --categ "Food"
  expense-tracker -list --month 8
  expense-tracker -sum 
  expense-tracker -sum --month 8
  expense-tracker -sum --cat

options:`
		fmt.Fprintln(os.Stderr, usage)
		fs.PrintDefaults()
	}

	// define flags
	fs.BoolVar(&flags.Add, "add", false, "Add a new expense")
	fs.BoolVar(&flags.List, "list", false, "List all expense")
	fs.BoolVar(&flags.Summary, "sum", false, "Get total expense summary")
	fs.BoolVar(&flags.Cate, "cat", false, "Get all Category")

	fs.IntVar(&flags.Update, "upd", NoIDSelected, "Update expense by ID")
	fs.IntVar(&flags.Delete, "del", NoIDSelected, "Delete expense by ID")
	fs.IntVar(&flags.MonthID, "month", NoIDSelected, "month number")

	fs.Float64Var(&flags.Amount, "amount", NoIDSelected, "amount of expense")

	fs.StringVar(&flags.Description, "desc", "", "description of expense")
	fs.StringVar(&flags.Category, "categ", "", "category of expense")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	return flags, nil
}

// DetermineCommand determines which command to execute based on flags
func (f *CmdFlags) DetermineCommand() (string, error) {
	commandCount := 0
	var command string

	if f.Add {
		commandCount++
		command = "add"
	}

	if f.List {
		commandCount++
		command = "list"
	}

	if f.Summary {
		commandCount++
		command = "summary"
	}

	if f.Update != NoIDSelected {
		commandCount++
		command = "update"
	}

	if f.Delete != NoIDSelected {
		commandCount++
		command = "delete"
	}

	if commandCount == 0 {
		return "", fmt.Errorf("no command specified. Use -h for help")
	}

	if commandCount > 1 {
		return "", fmt.Errorf("multiple commands specified. Use only one command at a time")
	}

	return command, nil
}
