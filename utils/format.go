package utils

import "strconv"

func MoneyFormatter(amount int) string {
	amountStr := strconv.Itoa(amount)
	n := len(amountStr)

	result := ""
	count := 0

	for i := n - 1; i >= 0; i-- {
		result = string(amountStr[i]) + result
		count++

		if count%3 == 0 {
			result = "," + result
		}
	}

	result = "Rp" + result + ".00"

	return result
}