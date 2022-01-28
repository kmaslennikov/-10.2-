package main

import (
	"fmt";
  "math"
)

func main() {
	fmt.Println("Введите денежную сумму:")
	var deposit float64
	fmt.Scan(&deposit)

	fmt.Println("Введите ежемесячный процент капитализации:")
	var percent float64
	fmt.Scan(&percent)

	fmt.Println("Введите срок вклада(количество лет):")
	var period int
	fmt.Scan(&period)

	months := period * 12
	result := deposit
  benefit:=0.0
	for i := 1; i <= months; i++ {
		percentSum := result * percent * 0.01
    rouding:= math.Floor(percentSum*100)/100
    benefit += percentSum - rouding
  
		result += rouding
	}
fmt.Println(result)
fmt.Println(benefit)
}
