package main

import "fmt"

func main() {

	var confranceName = "SG Turf"
	const conferenceTicket = 50
	var availabelTickets uint = 50

	availabelTickets = -1

	fmt.Println("Welcome To", confranceName, "Booking Application")
	fmt.Println("Total Tickets Available:", availabelTickets)
	fmt.Println("Ticket Price: $", conferenceTicket)

	var userName string

	var userTicket = 2
	userName = "tom"

	fmt.Println("Hello", userName, "You have booked", userTicket, "tickets")

}
