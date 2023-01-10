package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetusers()

	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("")

	//user input from here
	firstName, lastName, email, userTickets := getUserInput()

	//validate user input
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		//booking the application
		bookTicket(userTickets, firstName, lastName, email)
		//sending the email
		wg.Add(1) //goroutines to wait
		go sendTicket(userTickets, firstName, lastName, email)

		//call function print name
		firstNames := getFirstNames()
		fmt.Printf("The firstNames of our bookings are : %v \n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out.Come back next year.")
			//break
		}

	} else {

		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered is does not contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is inValid")
		}

		fmt.Printf("Your input data is invalid, try again\n")

	}
	wg.Wait()
}

func greetusers() {
	fmt.Printf("Welcome to our %v booking application", conferenceName)
	fmt.Println("")
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Println("")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	//fmt.Printf("The firstNames of our bookings are : %v \n", firstNames)
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask for their name
	println("Enter your firstName:")
	fmt.Scan(&firstName)

	println("Enter your last\tName:")
	fmt.Scan(&lastName)

	println("Enter your email:")
	fmt.Scan(&email)

	println("Enter the Number of Tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	// fmt.Printf("The whole slice : %v \n", bookings)
	// fmt.Printf("The first value slice : %v \n", bookings[0])
	// fmt.Printf("Type fo slice : %T \n", bookings)
	// fmt.Printf("slice length : %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("***************************")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("***************************")
	wg.Done()
}
