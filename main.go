package main

import (
	// "bufio"
	// "fmt"
	// "income_splitter/logic"
	// "income_splitter/utils"
	// "os"
	// "strconv"
	// "strings"
	// "time"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	router.Run("localhost:8080")

	// option := ""
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Print("User: ")
	// user,_ := reader.ReadString('\n')
	// user = strings.TrimSpace(user)

	// fmt.Print("\n")

	// fmt.Printf("Hello, %s\n", user)



	// for option != "2"{
	// 	fmt.Println("What would you like to do: ")
	// 	fmt.Println("1. Split money")
	// 	fmt.Println("2. Exit")

	// 	option, _ := reader.ReadString('\n')
	// 	option = strings.TrimSpace(option)
	// 		switch option{
	// 			case "1":
	// 				fmt.Print("Please input your income in rupiah (do not use dots or comma): ")
	// 				incomeReader := bufio.NewReader(os.Stdin)
	// 				incomeStr,_ := incomeReader.ReadString('\n')
	// 				incomeStr = strings.TrimSpace(incomeStr)
					
	// 				income, err := strconv.Atoi(incomeStr)
	// 				if err != nil{
	// 					fmt.Println("Invalid number please input numbers without any text")
	// 				}

	// 				divide := logic.Divider(income)
	// 				fmt.Printf("Needs: %s\n", utils.MoneyFormatter(divide.NeedsExpense))
	// 				fmt.Printf("Savings: %s\n", utils.MoneyFormatter(divide.SavingExpense))
	// 				fmt.Printf("Independents: %s\n", utils.MoneyFormatter(divide.IndependetMoney))

	// 			case "2":
	// 				fmt.Println("Goodbye!")
	// 				time.Sleep(2 * time.Second)
	// 				return
	// 			default:
	// 				fmt.Println("Invalid options")
	// 		}		
	// }	
}