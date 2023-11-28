package main

import (
	"bookking-application/shared"
	"fmt"
	"time"
	"sync"
	
	
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = make([]userData, 0)
type userData struct{
	firstName string
	lastName string
    email string
	userTickets uint
}

var wg = sync.WaitGroup{}
func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := shared.Vallidateuserinput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendticket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("all tickets for our conference are sold out, do chek our website for next events")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered  should contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
	wg.Wait()
}

func printFirstNames() []string {
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

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}




func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	 
	// create map for users
	
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,

	}
	

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
func sendticket(userTickets uint, firstName string, lastName string, email string)  {
	time.Sleep(6 * time.Second)
	var ticket = fmt.Sprintf("%v tickets %v %v for the user", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("sending ttickets to %v email adress  %v\n ",ticket, email )
	fmt.Println("#####################")
	wg.Done()
}
