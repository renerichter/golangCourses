package main

import "fmt"

func getExpenseReport(e expense) (string, string, float64) {
	em, ok := e.(email)
	if ok {
		return "email", em.toAddress, em.cost()
	}
	s, ok := e.(sms)
	if ok {
		return "sms", s.toPhoneNumber, s.cost()
	}
	return "invalid", "", 0.0
}

func getExpenseReportSwitched(e expense) (string, string, float64) {
	switch v := e.(type) {
	case email:
		return "email", v.toAddress, v.cost()
	case sms:
		return "sms", v.toPhoneNumber, v.cost()
	default:
		return "invalid", "", 0.0
	}
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}
type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}
type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}
type invalid struct {
}

func test03(e expense) {
	// eType, toAddress, cost := getExpenseReport(e)
	eType, toAddress, cost := getExpenseReportSwitched(e)
	fmt.Printf("Your %v to %v costs %v.\n", eType, toAddress, cost)
	fmt.Println("===========================")
}

func main() {
	test03(email{
		isSubscribed: true,
		body:         "blabla blub blu",
		toAddress:    "123@gmail.com",
	})
	test03(sms{
		isSubscribed:  false,
		body:          "Hi you.",
		toPhoneNumber: "01489652148",
	})

}
