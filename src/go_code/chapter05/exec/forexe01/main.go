package main

import "fmt"

/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-14 16:34:10
 * @LastEditTime: 2021-11-14 20:48:58
 * @LastEditors: Do not edit
 */
func main() {
	// // for j := 1; j <= 5; j++ {
	// // 	var sum float64 = 0
	// // 	for i := 1; i <= 5; i++ {
	// // 		var score float64
	// // 		fmt.Printf("请输入第 %d 个班级,第 %d 个学生的成绩\n", j, i)
	// // 		fmt.Scanln(&score)
	// // 		sum += score
	// // 	}
	// // 	fmt.Printf("%d班级的平均分为:%v\n", j, sum/5)
	// var n int = 10
	// for i := 1; i <= n; i++ {
	// 	for k := 1; k <= n-i; k++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for j := 1; j <= 2*i-1; j++ {
	// 		if j == 1 || j == 2*i-1 || i == n {
	// 			fmt.Print("*")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}

	// 	}
	// 	fmt.Println()
	// }
	// for i := 1; i <= 9; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d * %d = %d\t", j, i, i*j)
	// 	}
	// 	fmt.Println()
	// }
	var year int = 1999
	var month int = 2
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(31)
	case 4, 6, 9, 11:
		fmt.Println(30)
	default:
		if ((year%4 == 0 && year%100 != 0) || year%400 == 0) && month == 2 {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	}

}
