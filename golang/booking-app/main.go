package main
import (
	"fmt"
	// "strings"
)

func main(){
	conferenceName := "Go conference"

	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets) 
	fmt.Println("Get your tickets here to attend")

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


		var bookings []string


		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])

		fmt.Printf("Thank you %v %v for booking %v tickets. you will receive confirmation at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remain for %v\n", remainingTickets, conferenceName)

		// firstNames := []string{}
		// for index, booking := range bookings {
		// 	var names = strings.Fields(booking)
		// 	var firstName = names[0]
		// 	firstNames = append(firstNma)
		// }
		fmt.Printf("These are all our bookings: %v\n", bookings)
	}

}
