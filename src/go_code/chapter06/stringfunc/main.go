/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2022-01-09 00:23:19
 * @LastEditTime: 2022-01-09 23:31:07
 * @LastEditors: Do not edit
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// str1 := "hello康楠"
	// fmt.Println("字符串的长度:", len(str1))
	// str2 := "hello康楠"
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("字符=%c\n", str2[i])
	// }
	// //转rune切片
	// str3 := []rune(str2)
	// for i := 0; i < len(str3); i++ {
	// 	fmt.Printf("字符=%c\n", str3[i])
	// }
	//3)字符串转整数: n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误: ", err)
	} else {
		fmt.Println("转换结果: ", n)
	}
	str := strconv.Itoa(12345)
	fmt.Println("str=", str)
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)
	//[]byte转字符串:str=string([]byte{97,98,99})
	str2 := string([]byte{97, 98, 99})
	fmt.Printf("bytes=%v\n", str2)
	//10进制转 2,8,16进制: str = strconv.FomatInt(123, 2) // 2-> 8 , 16
	sstr := strconv.FormatInt(123, 2) // 2-> 8 , 16
	fmt.Printf("123对应的二进制:%v\n", sstr)
	ssstr := strconv.FormatInt(123, 16) // 2-> 8 , 16
	fmt.Printf("123对应的十六进制%v\n", ssstr)
	//查找子串是否在指定的字符串中: strings.Contains("seafood","foo") /true
	b := strings.Contains("seafood", "foo")
	fmt.Printf("seafood中是否含有子串foo  %v\n", b)
	//9)统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "e") //4 .
	d := strings.Count("ceheese", "e")
	fmt.Printf("ceheese中含有%v个e\n", d)
	// 不区分大小写的字符串比较(==是区分字母大小写的）:
	// fmt.Println(strings.EqualFold("abc","Abc"))
	fmt.Println(strings.EqualFold("abc", "Abc"))
	fmt.Println("abc" == "Abc")
	//返回子串在字符串第一次出现的 ndex值，如果没有返回- 1：strings.Index（“NLT_abc”，“abc”）
	s := strings.Index("NLT_abc", "abc")
	fmt.Printf("index=%v\n", s)
	s = strings.LastIndex("go golang", "go")
	fmt.Printf("index=%v\n", s)
	//将指定的子串替换成 另外一个子串：
	//strings.Replace(" 2 0 " , 1 //n可以指定你希望替换几个，如果n=-1表示全部替换
	str = strings.Replace("go go hello", "go", "go语言", -1)
	fmt.Println(str)
	//14) 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：
	strs := strings.Split("hello,wrold,ok", ",")
	fmt.Printf("strs=%v\n", strs)
	//将字符串的字母进行大小写的转换:
	// strings.ToLower("Go )   //go strings.ToUpper("G // GO
	m := strings.ToLower("Go")
	fmt.Printf("Upper=%v\n", m)
	m = strings.ToUpper("Go")
	fmt.Printf("Upper=%v\n", m)
	//16) 将字符串左右两边的空格去掉：
	//strings.TrimSpace(" tn a lone gopher ntrn")
	m = strings.TrimSpace(" tn a lone gopher ntrn   ")
	fmt.Printf("trimspace=%v\n", m)
	//17)将字符串左右两边指定的字符去掉 :
	// strings.Trim("! hello! "," !") //["hello"]//将左右两边 ! “”去掉
	m = strings.Trim("! hello! ", " !") //["hello"]
	fmt.Printf("trim=%v\n", m)
	//20) 判断字符串是否以指定的字符串开头:
	//strings.HasPrefix("ftp://192.168.10.1","ftp") // true
	fmt.Println(strings.HasPrefix("ftp://192.168.10.1", "ftp"))
	fmt.Println(strings.HasSuffix("ftp://192.168.10.1", "ftp"))
}
