package chapter12

import "fmt"

func main() {
	//声明一个变量,保存接收用户输入的选项
	key := ""
	//声明一个变量,控制是否退出 for 循环
	loop := true
	//定义账户的余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//收支的详情使用字符串来记录
	details := "收支\t账户金额\t收支金额\t说     明\n"
	//显示主菜单
	for {
		fmt.Println("--------------------家庭收支记账软件--------------------")
		fmt.Println("                     1. 收支明细                      ")
		fmt.Println("                     2. 登记收入                      ")
		fmt.Println("                     3. 登记支出                      ")
		fmt.Println("                     4. 退出软件                      ")
		fmt.Println("-----------------------------------------------------")
		fmt.Println("请选择(1-4): ")

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("\n--------------------当前收支明细记录--------------------")
			fmt.Println(details)
		case "2":
			fmt.Println("\n--------------------正在进行登记收入--------------------")
			fmt.Println(" 本次收入金额: ")
			fmt.Scanln(&money)
			balance += money
			fmt.Println(" 本次收入说明: ")
			fmt.Scanln(&note)
			//将收入情况,拼接到 details 变量
			details += fmt.Sprintf("收入\t%v\t\t%v\t\t%v\n", balance, money, note)
		case "3":
			fmt.Println("\n--------------------正在进行登记支出--------------------")
			fmt.Println(" 本次支出金额: ")
			fmt.Scanln(&money)
			//这里需要做一个必要的判断
			if money > balance {
				fmt.Println("支出的余额不足!!!")
				break
			}
			balance -= money
			fmt.Println(" 本次支出说明: ")
			fmt.Scanln(&note)
			//将收入情况,拼接到 details 变量
			details += fmt.Sprintf("支出\t%v\t\t%v\t\t%v\n", balance, money, note)
		case "4":
			loop = false
		default:
			fmt.Println("--------------------请输入正确的选项--------------------")
		}
		if !loop {
			fmt.Println("正在退出家庭记账软件的使用")
			break
		}
	}
}
