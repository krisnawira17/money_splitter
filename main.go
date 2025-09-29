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

func divider(income int) {
	result := &Split{
		NeedsExpense: income * (50/100),
		SavingExpense: NeedsExpense * (30/100),
		IndependetMoney: SavingExpense * (20/100),
	}
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
	
	income, err := strconv.Atoi(income)
	if err != nil{
		fmt.Println("Invalid number:", err)
		return
	}

	divider(income)
	

	// reader := bufio.NewReader(os.Stdin)
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	// fmt.Println("Hello", name)
}