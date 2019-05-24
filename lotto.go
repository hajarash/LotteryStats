package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
const SIZE_LOTTO_MAX = 8
const SIZE_DAILY_GRAND = 6
const LOTTO_MAX_CODE = 1
const DAILY_GRAND_CODE = 2
const MAX_LOTTO_49 = 49
const LOTTO_MAX_PRICE = 5
const DAILY_GRAND_PRICE = 3


var max_budget  = 3000

var lottoCode int

func main() {
	fmt.Print("1) lotto max\n2) lotto grand\n")
	var i int
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		os.Exit(0)
	}
	fmt.Print("budget (positive):")

	var d int
	_, er := fmt.Scanf("%d", &d)

	if er != nil {
		os.Exit(0)
	}
	if d > 0{
		max_budget = d
	}
	fmt.Printf("running the algorithm for $%d budget\n", max_budget)

	lottoCode = i
	if lottoCode != 1 && lottoCode!= 2{
		main()
	}else {
		testRun(lottoCode)
	}
}


func testRun(lottoType int) {
	switch lottoType {
	case LOTTO_MAX_CODE:
		fmt.Print("Enter winning numbers comma splitted (7 numbers and 1 bonus):\n")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		myList := strings.Split(text, ",")
		var numbers []int
		for i := 0; i < SIZE_LOTTO_MAX; i++ {
			num, err := strconv.ParseInt(strings.Split(myList[i], "\n")[0],10,64)
			if err != nil {
				print(err.Error())
				os.Exit(0)
			}
			numbers = append(numbers, int(num))
		}
		examineNumbers(numbers)
		break
	case DAILY_GRAND_CODE:
		fmt.Print("Enter winning numbers comma splitted (5 numbers and 1 GRAND number):\n")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		myList := strings.Split(text, ",")
		var numbers []int
		for i := 0; i < SIZE_DAILY_GRAND; i++ {
			num, err := strconv.ParseInt(strings.Split(myList[i], "\n")[0],10,64)
			if err != nil {
				print(err.Error())
				os.Exit(0)
			}
			numbers = append(numbers, int(num))
		}
		examineNumbers(numbers)
		break
	}

}
