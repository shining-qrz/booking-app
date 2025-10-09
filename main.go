package main

import "fmt"

func main() {
	conferenceName := "Go Conference" //等价于 var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	var username string
	var ticketCount uint
	var email string

	// 定义数组并初始化
	// var bookings = [50]string {"张三", "李四", "王五"}

	// 语法糖：定义空数组
	// var bookings [50]string

	var bookings []string

	for {
		fmt.Printf("Welcome to %v booking application!\n", conferenceName)
		fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

		fmt.Println("请输入用户名：")
		fmt.Scan(&username)

		fmt.Println("请输入需预定票的数量：")
		fmt.Scan(&ticketCount)

		fmt.Println("请输入您的邮箱：")
		fmt.Scan(&email)

		if remainingTickets-ticketCount <= 50 {
			remainingTickets = remainingTickets - ticketCount
			bookings = append(bookings, username)

			fmt.Printf("恭喜，%v！您已成功预定%v张票，我们稍后将会发送邮件至：%v\n", username, ticketCount, email)

			fmt.Println("剩余票数：", remainingTickets)
		} else {
			fmt.Printf("剩余票数不足，无法预定！\n")
		}
	}

}
