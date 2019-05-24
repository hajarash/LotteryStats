package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func examineNumbers(numbers []int) {
	switch lottoCode {
		case LOTTO_MAX_CODE:
		 	examineLottoMax(numbers)
		 	break
		case DAILY_GRAND_CODE:
			examineDailyGrand(numbers)
			break
	}
}

func generateEmptyPool() []bool {
	var pool []bool

	for i := 0; i < MAX_LOTTO_49; i++ {
		pool = append(pool,false)
	}
	if (lottoCode == LOTTO_MAX_CODE) {
		pool = append(pool,false)
	}
	return pool
}

func getNumberOfUnusedNumbers(pool []bool) int {
	count :=0
	for _, element:= range pool {
		if !element {
			count ++
		}
	}

	return count
}

func confrimNumber(pool []bool, num int) int {
	isTrue := true
	for isTrue {
		isTrue = false
		if pool[num]{
			num++
			isTrue = true
		}
	}
	return num
}

func adjustNumber(pool []bool, num int) int {
	adjustedNumber := num
	for i:= 0 ; i <= adjustedNumber; i++ {
		if pool[i] {
			adjustedNumber++
		}
	}
	return confrimNumber(pool, num)
}

func GenerateNumber(pool []bool) ([]bool, []int){
	availableNumbers := getNumberOfUnusedNumbers(pool)
	if availableNumbers < SIZE_DAILY_GRAND* 2{
		pool = generateEmptyPool()
		return GenerateNumber(pool)
	} else{
		var draw []int
		if lottoCode == LOTTO_MAX_CODE {
			for i := 0; i< SIZE_LOTTO_MAX-1; i++ {
				num := rand.Intn(availableNumbers)
				num = adjustNumber(pool, num)
				pool[num] = true
				availableNumbers--
				draw = append(draw, num+1)
			}
			return pool, draw
		} else {
			for i := 0; i< SIZE_DAILY_GRAND-1; i++ {
				num := rand.Intn(availableNumbers)
				num = adjustNumber(pool, num)
				pool[num] = true
				availableNumbers--
				draw = append(draw, num+1)
			}
			return pool, draw
		}
	}
}

func checkForMatch(numbers[] int,number int ,startingIndex int) (bool, int) {
	done :=false
	match := false
	for ;!done && startingIndex < len(numbers)-1; startingIndex++ {
		if number == numbers[startingIndex] {
			done = true
			match = true
		}else if number < numbers[startingIndex] {
			done = true
		}
	}
	startingIndex--
	return match, startingIndex
}

func checkForWin(numbers [] int, draw [] int) bool {
	bonus := numbers[len(numbers)-1]
	bonus= bonus
	startingIndex := 0
	numOfMatches:= 0
	for _,element:=range draw {
		var match bool
		match, startingIndex = checkForMatch(numbers, element, startingIndex)
		if match {
			numOfMatches++
		}
	}
	return numOfMatches >= 6
}

func examineLottoMax(numbers []int) {
	pool := generateEmptyPool()
	x := max_budget
	win := false
	for x > 0 && !win {
		var draw []int
		x-=LOTTO_MAX_PRICE
		pool, draw = GenerateNumber(pool)
		sort.Ints(draw)
		win = checkForWin(numbers, draw)
		for _,element:=range draw{
			print (element)
		}
		println()
	}
	if win {
		amount_spent := max_budget-x
		print ("WON! ")
		fmt.Printf("money spent: %d", amount_spent)
	} else {
		fmt.Printf("you just lost %d dollars", max_budget)
	}
}

func examineDailyGrand(numbers []int) {
	pool := generateEmptyPool()
	x := max_budget
	win := false
	for x > 0 && !win {
		var draw []int
		x-=DAILY_GRAND_PRICE
		pool, draw = GenerateNumber(pool)
		sort.Ints(draw)
		win = checkForWin(numbers, draw)
	}
	if win {
		amount_spent := max_budget-x
		print ("WON! ")
		fmt.Printf("money spent: %d", amount_spent)
	} else {
		fmt.Printf("you just lost %d dollars", max_budget)
	}
}

