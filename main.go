package main

import "fmt"

func main() {

	var confranceName = "SG Turf"
	const conferenceTicket = 50
	var availabelTickets = 50

	fmt.Println("Welcome To", confranceName, "Booking Application")
	fmt.Println("Total Tickets Available:", availabelTickets)

	// Print the conference ticket
	fmt.Println("Ticket Price: $", conferenceTicket)

	fmt.Print(conferenceTicket)
}
