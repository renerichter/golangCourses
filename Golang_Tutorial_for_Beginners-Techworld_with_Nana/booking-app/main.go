package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = conferenceTickets

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup()

var bookings = make([]UserData, 0)

func main() {
	greetUsers()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, userTickets, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("These are our bookings: %v", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is sold out. Come back next year.")
				break
			} else {
				fmt.Printf("%v tickets are still remaining\n", remainingTickets)
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email entered wrong.")
			}
			if !isValidTicketNumber {
				fmt.Println("Wrong ticket number entered.")
			}
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName) //same as below
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name:       ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name:        ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email address:    ")
	fmt.Scan(&email)
	fmt.Printf("How many tickets do you want?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {
	remainingTickets -= userTickets

	// create a map for a user
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. We will send a confirmation email to %v.\n", firstName, lastName, userTickets, email)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %v \n to email %v", ticket, email)
	fmt.Println("######################")
	wg.Done()
}
