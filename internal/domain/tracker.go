package domain

import (
	"sync"
)

const NoIDSelected = -1

// -----------------------
// Expense Tracker
// -----------------------
type ExpenseTracker struct {
	Expenses []*Expense      `json:"Expenses"`
	Budgets  map[int]float64 `json:"Budgets"`
	NextID   int
}

var (
	instance *ExpenseTracker
	once     sync.Once
)

func GetExpenseTracker() *ExpenseTracker {
	once.Do(func() {
		instance = &ExpenseTracker{
			Expenses: make([]*Expense, 0),
			Budgets: map[int]float64{
				1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0,
				7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0,
			},
			NextID: 0,
		}
	})

	return instance
}

// add
func (t *ExpenseTracker) Add(amount float64, desc, cate string) {
	t.Expenses = append(t.Expenses, newExpense(t.NextID, amount, desc, cate))
}

// delete
func (t *ExpenseTracker) Delete(index int) {
	t.Expenses = append(t.Expenses[:index], t.Expenses[index+1:]...)
}

// update
func (t *ExpenseTracker) Update(index int, amount float64, desc, cate string) {
	currentExpense := t.Expenses[index]

	if amount != float64(NoIDSelected) {
		currentExpense.Amount = amount
	}

	if desc != "" {
		currentExpense.Description = desc
	}

	if cate != "" {
		currentExpense.Category = cate
	}

}

// summary
func (t *ExpenseTracker) Summary(expeneslist []*Expense) float64 {
	var totalAmount float64

	for _, expense := range expeneslist {
		totalAmount += expense.Amount
	}

	return totalAmount
}

// Set Budgets for all month
func (t *ExpenseTracker) SetBudget(monthID int, amount float64) {
	if _, ok := t.Budgets[monthID]; !ok {
		t.Budgets[monthID] = amount
	}
}

// Get budget of the month
func (t *ExpenseTracker) GetBudgetofTheMonth(monthID int) float64 {
	if val, ok := t.Budgets[monthID]; ok {
		return val
	}

	return 0
}

// update NextID
func (t *ExpenseTracker) UpdateNextID() int {
	maxId := 0

	for _, expense := range t.Expenses {
		if expense.Id > maxId {
			maxId = expense.Id
		}
	}

	return maxId + 1
}
