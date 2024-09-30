package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return cost + costSpouse, nil
}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v char", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func test01(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("==================")
	fmt.Println("Msg for Cust: ", msgToCustomer)
	fmt.Println("Msg for Spouse: ", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("Total cost %v.\n", totalCost)
}

func main() {
	test01("hi you", "and you, too")
	test01("hi you", "uiaenrduinaerduianeduiarenduirtaenuitaen")
}
