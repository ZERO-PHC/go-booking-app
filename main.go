package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets int = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and  %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		//ask user for their name
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email", firstName, lastName, userTickets)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			// range provides the index and value for each element
			// range can iterate over elements with different data structures (not only arrs and slices)

			// _ Blank identifier
			for _, booking := range bookings {
				// splits the string with white space as separator
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our Conference is booked out. Come back nex year")
				break
			}

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
		} else {
			fmt.Printf("We only have %v tickets remaining", remainingTickets)
		}
		//implicit continue to next iteration
	}

}
