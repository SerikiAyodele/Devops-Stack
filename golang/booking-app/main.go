package main
import (
	"fmt"
	"strings"
)

var conferenceName = "Go conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = []string{}

func main(){
	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstnames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Booked out")
				break
			} 
		} else {
			if !isValidName {
				fmt.Println("first or last name you entered was too short")
			}
			if !isValidEmail {
				fmt.Println("email you entered does not contain the @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
	  } 
	}
}

func greetUsers(){
	fmt.Printf("Welcome to the %v app\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets) 
	fmt.Println("Get your tickets here to attend")
}

func getFirstnames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("enter your firstname:")
	fmt.Scan(&firstName)
	fmt.Println("enter your lastname:")
	fmt.Scan(&lastName)
	fmt.Println("enter your email:")
	fmt.Scan(&email)
	fmt.Println("enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remain for %v\n", remainingTickets, conferenceName)

	return remainingTickets, bookings
}