package common

import (
	"fmt"
	"strings"
	"time"
)

var ConferenceName = "qrz Conference"

const ConferenceTickets = 50

var RemainingTickets uint = 50
var Bookings []User

type User struct {
	firstName   string
	lastName    string
	userName    string
	email       string
	ticketCount uint
}

// Greeting 输出招呼语
func Greeting() {
	fmt.Printf("Welcome to %v booking application!\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", ConferenceTickets, RemainingTickets)
}

// GetUserInput 获取用户输入
func GetUserInput() User {
	var user User
	//读取用户输入
	fmt.Println("请输入姓氏：")
	fmt.Scan(&user.firstName)
	fmt.Println("请输入名字：")
	fmt.Scan(&user.lastName)
	fmt.Println("请输入需预定票的数量：")
	fmt.Scan(&user.ticketCount)
	fmt.Println("请输入您的邮箱：")
	fmt.Scan(&user.email)

	user.userName = user.firstName + user.lastName

	return user
}

// ValidateUserInput 验证用户输入
func ValidateUserInput(user User) (bool, bool, bool) {
	isValidName := len(user.firstName) >= 2 && len(user.lastName) >= 2
	isValidEmail := strings.Contains(user.email, "@")
	isValidTicketCount := user.ticketCount > 0 && user.ticketCount <= RemainingTickets

	return isValidName, isValidEmail, isValidTicketCount
}

// BookTickets 预定票
func BookTickets(user User) {
	RemainingTickets = RemainingTickets - user.ticketCount
	Bookings = append(Bookings, user)
}

// OutPutBookingInfo 输出预定信息
func OutPutBookingInfo(user User) {
	fmt.Printf("恭喜，%v！您已成功预定%v张票，我们稍后将会发送邮件至：%v\n", user.userName, user.ticketCount, user.email)
	fmt.Println("剩余票数：", RemainingTickets)
	var bookingNames []string
	for _, booking := range Bookings {
		bookingNames = append(bookingNames, booking.userName)
	}
	fmt.Printf("当前预定成功者：%v\n", bookingNames)
}

func OutPutEndLine() {
	fmt.Println("\n---------------------------------------------------")
	time.Sleep(5 * time.Second)
	fmt.Println()
}
