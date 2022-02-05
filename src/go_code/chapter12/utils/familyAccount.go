package utils

import "fmt"

type FamilyAccount struct {
	//声明一个变量,保存接收用户输入的选项
	key string
	//声明一个变量,控制是否退出 for 循环
	loop bool
	//定义账户的余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//收支的详情使用字符串来记录
	details string
}

//给结构体绑定方法
//显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("--------------------家庭收支记账软件--------------------")
		fmt.Println("                     1. 收支明细                      ")
		fmt.Println("                     2. 登记收入                      ")
		fmt.Println("                     3. 登记支出                      ")
		fmt.Println("                     4. 退出软件                      ")
		fmt.Println("-----------------------------------------------------")
		fmt.Println("请选择(1-4): ")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.ShowDetails()
		case "2":
			this.Income()
		case "3":
			this.Pay()
		case "4":
			this.Exit()
		default:
			fmt.Println("--------------------请输入正确的选项--------------------")
		}
		if !this.loop {
			fmt.Println("正在退出家庭记账软件的使用")
			break
		}
	}
}
func (this *FamilyAccount) ShowDetails() {
	fmt.Println("\n--------------------当前收支明细记录--------------------")
	fmt.Println(this.details)
}

func (this *FamilyAccount) Income() {
	fmt.Println("\n--------------------正在进行登记收入--------------------")
	fmt.Println(" 本次收入金额: ")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println(" 本次收入说明: ")
	fmt.Scanln(&this.note)
	//将收入情况,拼接到 details 变量
	this.details += fmt.Sprintf("收入\t%v\t\t%v\t\t%v\n", this.balance, this.money, this.note)
}

func (this *FamilyAccount) Pay() {
	fmt.Println("\n--------------------正在进行登记支出--------------------")
	fmt.Println(" 本次支出金额: ")
	fmt.Scanln(&this.money)
	//这里需要做一个必要的判断
	if this.money > this.balance {
		fmt.Println("支出的余额不足!!!")
		return
	}
	this.balance -= this.money
	fmt.Println(" 本次支出说明: ")
	fmt.Scanln(&this.note)
	//将收入情况,拼接到 details 变量
	this.details += fmt.Sprintf("支出\t%v\t\t%v\t\t%v\n", this.balance, this.money, this.note)
}
func (this *FamilyAccount) Exit() {
	this.loop = false
}
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000,
		money:   0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说     明\n",
	}
}
