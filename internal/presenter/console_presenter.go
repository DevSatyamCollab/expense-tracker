package presenter

import (
	"expense-tracker/internal/domain"
	"expense-tracker/internal/service"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

const dateFormat string = "02-01-2006"

// ---------------------
// console presenter
// ---------------------
type ConsolePresenter struct {
	service *service.ExpenseService
}

func NewConsolePresenter(s *service.ExpenseService) *ConsolePresenter {
	return &ConsolePresenter{service: s}
}

// show success message
func (p *ConsolePresenter) Success(msg string) {
	fmt.Println(msg)
}

func (p *ConsolePresenter) ShowTotalSummary() {
	price := strconv.FormatFloat(p.service.GetTotalSummary(), 'f', 2, 64)
	fmt.Println("Total expenses: ", price)
}

func (p *ConsolePresenter) ShowSummaryofMonth(monthID int) {
	price := p.service.GetSummaryOfMonth(monthID)
	month := time.Month(monthID).String()
	fmt.Printf("Total expenses for %s: %.2f\n", month, price)
}

func (p *ConsolePresenter) ShowSummaryofCategory() {
	categories := p.service.GetAllCategories()

	if len(categories) == 0 {
		fmt.Println("categories list is empty")
		return
	}

	fmt.Println("SUMMARY OF CATEGORY:")
	fmt.Println("--------------------------------------------------------------------")

	tble := table.New(os.Stdout)
	tble.SetHeaders("Category", "Price")

	for _, category := range categories {
		tble.AddRow(category, fmt.Sprintf("%.2f", p.service.GetSummaryByCategory(category)))
	}

	tble.Render()
}

func (p *ConsolePresenter) ShowList(expenseList []*domain.Expense) {
	if len(expenseList) == 0 {
		fmt.Println("list is empty")
		return
	}

	fmt.Println("Expense List:")
	fmt.Println("--------------------------------------------------------------------")

	tble := table.New(os.Stdout)
	tble.SetHeaders("ID", "Date", "Description", "Amount", "Category")

	for _, expense := range expenseList {
		tble.AddRow(strconv.Itoa(expense.Id), expense.Date.Format(dateFormat),
			expense.Description, strconv.FormatFloat(expense.Amount, 'f', 2, 64), expense.Category)
	}

	tble.Render()
}
