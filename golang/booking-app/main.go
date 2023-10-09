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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. you will receive confirmation at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remain for %v\n", remainingTickets, conferenceName)

			fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Booked out")
				break
			}
		}

		fmt.Printf("Max no of tickets you can book is %v, please try again\n", remainingTickets)
		continue
	}

}
