package main

import (
	"fmt"
	"time"
)
	
var conferenceName string = "Go Conference"  
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	noOfTickets uint
}

func main() {
	
	greetUsers()


	for {


		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := ValidateUser(firstName, lastName, email,  userTickets, remainingTickets)

		if  isValidName && isValidEmail && isValidTickets{
			
			bookTickets(firstName, lastName, userTickets, email)

			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()

			fmt.Printf("The first names of our bookings : %v\n", firstNames)


			if remainingTickets == 0 {
				// end program 
				fmt.Println("\nThe conference is fully book. Come back next year.")
				break
			} 
		} else {
			
			if !isValidName {
				fmt.Println("\nFirst name or last name you entered is too short.")
			} 
			if !isValidEmail {
				fmt.Println("Your email must contain @")
			} 
			if !isValidTickets {
				fmt.Printf("The number of tickets you are booking is invalid. We have %v remaining tickets available.\n", remainingTickets)
			}
		} 
	} 
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Book your tickets now!")
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

	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want to input: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint, email string) {
	remainingTickets -= userTickets

	var userData = UserData {
		firstName : firstName,
		lastName : lastName,
		email: email,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("\nList of booking data: %v\n", bookings)

	fmt.Printf("\n\nThank you %v %v for booking %v tickets. You will receive the confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for our %v\n\n", remainingTickets, conferenceName)


}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v for user %v %v", userTickets, firstName, lastName)
	fmt.Println("\n\n##############")
	fmt.Printf("Sending ticket:\n %v \nto email %v.\n", ticket, email)
	fmt.Println("##############")
}
