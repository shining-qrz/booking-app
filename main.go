package main

import "fmt"

func main() {
	conferenceName := "Go Conference" //等价于 var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
