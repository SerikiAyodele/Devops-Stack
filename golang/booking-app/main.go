package main
import (
	"fmt"
	"strings"
)

func main(){
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	getFirstnames(bookings)
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
	validateUserInput (firstName, lastName, email, userTickets, remainingTickets)

	for {
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


		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. you will receive confirmation at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remain for %v\n", remainingTickets, conferenceName)

			fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])

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

func greetUsers(confName string, confTickets int, remainTickets uint){
	fmt.Printf("Welcome to the %v app\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainTickets) 
	fmt.Println("Get your tickets here to attend")
}

func getFirstnames(book []string) []string{
	firstNames := []string{}
	for _, booking := range book {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput (firstName string, lastName string, email string, userTickets uint, remainingTickets uint){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
}