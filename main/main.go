package main

import (
	"booking-app/common"
	"fmt"
)

func main() {
	for {

		//输出招呼语
		common.Greeting()

		//获取用户输入
		user := common.GetUserInput()

		//校验用户输入
		isValidName, isValidEmail, isValidTicketCount := common.ValidateUserInput(user)
		if !isValidName {
			fmt.Println("您输入的姓名不合法，名字和姓氏至少需要2个字符，请重新输入。")
			common.OutPutEndLine()
			continue
		}
		if !isValidEmail {
			fmt.Println("您输入的邮箱不合法，邮箱中需要包含@符号，请重新输入。")
			common.OutPutEndLine()
			continue
		}
		if !isValidTicketCount {
			fmt.Printf("您输入的票数不合法，票数需要大于0且不能超过剩余票数，请重新输入。当前剩余票数：%v\n", common.RemainingTickets)
			common.OutPutEndLine()
			continue
		}

		//订票
		common.BookTickets(user)

		//输出订票信息
		common.OutPutBookingInfo(user)

		//发送会议票
		go common.SendEmail(user)

		//输出结束行
		common.OutPutEndLine()
	}
}
