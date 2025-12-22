package domain

import "time"

// --------------------
// Expense
// --------------------
type Expense struct {
	Id          int       `json:"ID"`
	Amount      float64   `json:"Amount"`
	Description string    `json:"Description"`
	Category    string    `json:"Category"`
	Date        time.Time `json:"Date"`
}

func newExpense(id int, amount float64, desc, cate string) *Expense {
	return &Expense{
		Id:          id,
		Amount:      amount,
		Description: desc,
		Category:    cate,
		Date:        time.Now(),
	}
}

func ValidateID(id int) error {
	if id < 0 {
		return ErrInvalidId
	}

	return nil
}

func ValidateAmount(amount float64) error {
	if amount < 0 {
		return ErrInvalidAmount
	}

	return nil
}

func ValidateMonthID(monthID int) error {
	if monthID < 1 || monthID > 12 {
		return ErrInvalidMonth
	}

	return nil
}

func ValidateDescription(desc string) error {
	if desc == "" {
		return ErrEmptyDescription
	}

	return nil
}

func ValidateCategory(category string) error {
	if category == "" {
		return ErrEmptyCategory
	}

	return nil
}

func ValidateExpense(amount float64, desc, cate string) error {
	var err error

	if err = ValidateAmount(amount); err != nil {
		return err
	}

	if err = ValidateDescription(desc); err != nil {
		return err
	}

	if err = ValidateCategory(cate); err != nil {
		return err
	}

	return nil
}
