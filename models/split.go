package models

type income struct {
	Id              string `json:"id"`
	NeedsExpense    int    `json:"needs_expense"`
	SavingExpense   int    `json:"saving_expense"`
	IndependetMoney int    `json:"independent_expense"`
}

func Divider(amount int) *income {
	result := &income{
		NeedsExpense:    amount * 50 / 100,
		SavingExpense:   amount * 30 / 100,
		IndependetMoney: amount * 20 / 100,
	}
	return result
}