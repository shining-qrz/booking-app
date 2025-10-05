package main

import "fmt"

func main() {
	conferenceName := "Go Conference" //等价于 var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	var username string
	var ticketCount uint
	var email string

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	fmt.Println("请输入用户名：")
	fmt.Scan(&username)

	fmt.Println("请输入需预定票的数量：")
	fmt.Scan(&ticketCount)

	fmt.Println("请输入您的邮箱：")
	fmt.Scan(&email)

	remainingTickets = remainingTickets - ticketCount

	fmt.Printf("恭喜，%v！您已成功预定%v张票，我们稍后将会发送邮件至：%v\n", username, ticketCount, email)

	fmt.Println("剩余票数：", remainingTickets)
}
