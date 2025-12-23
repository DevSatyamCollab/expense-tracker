package service

import (
	"expense-tracker/internal/domain"
	"expense-tracker/internal/storage"
	"fmt"
	"slices"
	"strings"
)

// -------------------------
// Expense Service
// -------------------------
type ExpenseService struct {
	tracker *domain.ExpenseTracker
	storage *storage.JsonStorage
}

func NewExpenseService(t *domain.ExpenseTracker, s *storage.JsonStorage) *ExpenseService {
	return &ExpenseService{tracker: t, storage: s}
}

// add expense and save
func (s *ExpenseService) AddExpense(amount float64, desc, cate string) error {
	s.tracker.Add(amount, desc, strings.ToLower(cate))

	if err := s.storage.Save(s.tracker); err != nil {
		return err
	}

	return nil
}

// update expense and save
func (s *ExpenseService) UpdateExpense(id int, amount float64, desc, cate string) error {
	index, err := s.findIndex(id)
	if err != nil {
		return fmt.Errorf("failed to update an expense: %w", err)
	}

	s.tracker.Update(index, amount, desc, cate)

	if err := s.storage.Save(s.tracker); err != nil {
		return err
	}

	return nil
}

// delete an expense and save
func (s *ExpenseService) DeleteExpense(id int) error {
	index, err := s.findIndex(id)
	if err != nil {
		return fmt.Errorf("failed to delete an expense: %w", err)
	}

	s.tracker.Delete(index)

	if err := s.storage.Save(s.tracker); err != nil {
		return err
	}

	return nil
}

// Get Total summary
func (s *ExpenseService) GetTotalSummary() float64 {
	return s.tracker.Summary(s.tracker.Expenses)
}

// Get summary of the month
func (s *ExpenseService) GetSummaryOfMonth(monthID int) float64 {
	return s.tracker.Summary(s.GetExpensesofMonth(monthID))
}

// Get summary of expenses by category
func (s *ExpenseService) GetSummaryByCategory(category string) float64 {
	return s.tracker.Summary(s.GetExpensesByCategory(category))
}

// Get all expenses
func (s *ExpenseService) GetAllExpenes() []*domain.Expense {
	return s.tracker.Expenses
}

// Get all categories
func (s *ExpenseService) GetAllCategories() []string {
	expenseList := s.tracker.Expenses
	seen := make(map[string]struct{})
	categoriesList := make([]string, 0)

	for _, expense := range expenseList {
		if _, ok := seen[expense.Category]; !ok {
			seen[expense.Category] = struct{}{}
			categoriesList = append(categoriesList, expense.Category)
		}
	}

	return categoriesList
}

// Get Expenses of the month
func (s *ExpenseService) GetExpensesofMonth(monthID int) []*domain.Expense {
	filter := make([]*domain.Expense, 0)

	for _, expense := range s.tracker.Expenses {
		if monthID == int(expense.Date.Month()) {
			filter = append(filter, expense)
		}
	}

	return filter
}

// Get Expenses by category
func (s *ExpenseService) GetExpensesByCategory(category string) []*domain.Expense {
	filter := make([]*domain.Expense, 0)

	for _, expense := range s.tracker.Expenses {
		if strings.EqualFold(category, expense.Category) {
			filter = append(filter, expense)
		}
	}

	return filter
}

// set budgets
func (s *ExpenseService) SetBudget(monthID int, amount float64) {
	s.tracker.SetBudget(monthID, amount)
}

// Get Budget
func (s *ExpenseService) GetBudget(monthID int) float64 {
	return s.tracker.GetBudgetofTheMonth(monthID)
}

// find expenseIndex
func (s *ExpenseService) findIndex(id int) (int, error) {
	index, found := slices.BinarySearchFunc(s.tracker.Expenses, id, func(expense *domain.Expense, targetID int) int {
		return expense.Id - targetID
	})

	if found {
		return index, nil
	}
	return -1, domain.ErrExpenseNotFound
}
