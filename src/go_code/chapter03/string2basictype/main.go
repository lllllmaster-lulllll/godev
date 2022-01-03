package main

import (
	"fmt"
	"strconv"
)

/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-07 16:37:31
 * @LastEditTime: 2021-11-07 17:28:49
 * @LastEditors: Do not edit
 */

func main() {
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v", b, b)

	var str2 string = "1234560"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type %T n1=%v", n1, n1)
}
