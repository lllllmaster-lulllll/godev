package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

type DivideError struct{
	dividee int
	divider int
}
func (de *DivideError) Error()string{
	strFormat:="Cannot proceed, the divider is zero.dividee:%ddivider:0"
	return fmt.Sprintf(strFormat,de.dividee)
}
func Divide(varDividee int, varDivider int)(result int, errorMsg string){
	if varDivider==0{
		dData:=DivideError{
			dividee:varDividee,
			divider:varDivider,
		}
		errorMsg=dData.Error()
		return
	}else{
		return varDividee/varDivider,""
	}
}

func main() {
	if result,errorMsg:=Divide(100,10);errorMsg==""{
		fmt.Println("100/10 = ", result)
	}
	if _,errorMsg:=Divide(100,0);errorMsg!=""{
		fmt.Println("errorMsg is: ",errorMsg)
	}
	str1 := ""    // 用于保存 "" 包裹的字符串
	str2 := ``    // 用于保存 `` 包裹的字符串
	timer := 1000 // 统一循环次数
	i := timer    // 计数器

	t1 := time.Now().UnixNano()
	for i > 0 {
		str1 += "测试\n测试\n测试\n测试\n测试\n测试\n"
		// str1 += "测试测试测试测试测试测试测试"
		i--
	}
	t2 := time.Now().UnixNano()
	println(`"" 所消耗的时间：`, t2-t1)

	i = timer // 重置计数器
	t3 := time.Now().UnixNano()
	for i > 0 {
		str2 += `测试` + "\n" + `测试` + "\n" + `测试` + "\n" +
			`测试` + "\n" + `测试` + "\n" + `测试` + "\n"
		// str2 += `测试测试测试测试测试测试测试`
		i--
	}
	t4 := time.Now().UnixNano()
	println("`` 所消耗的时间：", t4-t3)
	var chinese = "我是中国人， I am Chinese"
	fmt.Println("chinese length", len(chinese))
	fmt.Println("chinese word length", len([]rune(chinese)))
	fmt.Println("chinese word length", utf8.RuneCountInString(chinese))
}
