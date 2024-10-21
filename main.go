package main

import "fmt"

func main() {

	myname := "Ganesh"

	var confranceName = "GK Merta"
	const conferenceTicket = 50
	var availabelTickets uint = 50

	fmt.Println("Welcome To", confranceName, "Booking Application by", myname)
	fmt.Println("Total Tickets Available:", availabelTickets)
	fmt.Println("Ticket Price: $", conferenceTicket)

	var userName string
	fmt.Scan(&userName)
	fmt.Println("Hello", userName)

	fmt.Println("do you want to book the ticket 1=yes 2=no :")
	var yesno int
	fmt.Scan(&yesno)

	if yesno == 1 {
		fmt.Println("You have booked the ticket")
		availabelTickets--
	} else {
		fmt.Println("Thankyou for visiting the GK Mata")
	}

	fmt.Println("Hello", userName, "You have booked 1 tickets")

}
