package common

import (
	"fmt"
	"strings"
	"time"
)

var ConferenceName = "qrz Conference"

const ConferenceTickets = 50

var RemainingTickets uint = 50
var Bookings []string

// Greeting 输出招呼语
func Greeting() {
	fmt.Printf("Welcome to %v booking application!\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", ConferenceTickets, RemainingTickets)
}

// GetUserInput 获取用户输入
func GetUserInput() (string, string, string, uint, string) {
	var firstName string
	var lastName string
	var userName string
	var ticketCount uint
	var email string

	//读取用户输入
	fmt.Println("请输入姓氏：")
	fmt.Scan(&firstName)
	fmt.Println("请输入名字：")
	fmt.Scan(&lastName)
	fmt.Println("请输入需预定票的数量：")
	fmt.Scan(&ticketCount)
	fmt.Println("请输入您的邮箱：")
	fmt.Scan(&email)

	userName = firstName + lastName

	return firstName, lastName, userName, ticketCount, email
}

// ValidateUserInput 验证用户输入
func ValidateUserInput(firstName string, lastName string, email string, ticketCount uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := ticketCount > 0 && ticketCount <= RemainingTickets

	return isValidName, isValidEmail, isValidTicketCount
}

// BookTickets 预定票
func BookTickets(ticketCount uint, userName string) {
	RemainingTickets = RemainingTickets - ticketCount
	Bookings = append(Bookings, userName)
}

// OutPutBookingInfo 输出预定信息
func OutPutBookingInfo(userName string, ticketCount uint, email string) {
	fmt.Printf("恭喜，%v！您已成功预定%v张票，我们稍后将会发送邮件至：%v\n", userName, ticketCount, email)
	fmt.Println("剩余票数：", RemainingTickets)
	fmt.Printf("已预定用户名单：%v\n", Bookings)
}

func OutPutEndLine() {
	fmt.Println("\n---------------------------------------------------")
	time.Sleep(5 * time.Second)
	fmt.Println()
}
