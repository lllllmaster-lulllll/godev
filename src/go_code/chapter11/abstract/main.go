package main

import (
	"fmt"
)

//定义一个结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//方法
//存款
func (account *Account) Deposite(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("您输入的金额不正确")
		return
	}
	if money <= 0 {
		fmt.Println("您输入的金额不正确")

	}
	account.Balance += money
	fmt.Println("存款成功~~~")
}

//取款
func (account *Account) WithDraw(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("您输入的金额不正确")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("您输入的金额不正确")

	}
	account.Balance -= money
	fmt.Println("存款成功~~~")
}

//查询余额
func (account *Account) Query(pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("您输入的金额不正确")
		return
	}

	fmt.Printf("你的账号为:%v  余额为%v \n", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		AccountNo: "gs111111",
		Pwd:       "666",
		Balance:   100,
	}
	account.Query("666")

	fmt.Println()
}
