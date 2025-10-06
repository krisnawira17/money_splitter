package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Split struct{
	NeedsExpense int
	SavingExpense int
	IndependetMoney int
}

func divider(income int) *Split {	
	result := &Split{
		NeedsExpense: income * 50/100,
		SavingExpense: income * 30/100,
		IndependetMoney: income * 20/100,
	}

	return result
}

func moneyFormatter (amount int) string{
	amountStr := strconv.Itoa(amount)
	n := len(amountStr)

	result := ""
	count := 0
	
	for i := n-1; i >= 0; i--{
		result = string(amountStr[i]) + result
		count++
		
		if count % 3 == 0{
			result = "," + result
		}
	}

	result = "Rp" + result + ".00"

	return result
}

func main(){
	fmt.Print("User: ")
	userReader := bufio.NewReader(os.Stdin)
	user,_ := userReader.ReadString('\n')
	user = strings.TrimSpace(user)

	fmt.Print("\n")

	fmt.Printf("Hello, %s\n", user)

	fmt.Print("Please input your income in rupiah (do not use dots or comma): ")
	incomeReader := bufio.NewReader(os.Stdin)
	incomeStr,_ := incomeReader.ReadString('\n')
	incomeStr = strings.TrimSpace(incomeStr)
	
	income, err := strconv.Atoi(incomeStr)
	if err != nil{
		fmt.Println("Invalid number:", err)
		return
	}

	divide := divider(income)
	fmt.Printf("Needss: %s\n", moneyFormatter(divide.NeedsExpense) )
	fmt.Printf("Saving: %s\n", moneyFormatter(divide.SavingExpense) )
	fmt.Printf("Independent: %s\n", moneyFormatter(divide.IndependetMoney) )	
	
}