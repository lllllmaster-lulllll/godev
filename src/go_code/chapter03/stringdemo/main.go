/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-05 23:02:27
 * @LastEditTime: 2021-11-14 20:51:52
 * @LastEditors: Do not edit
 */
package main

import "fmt"

func main() {
	str3 := `
	/*
	* @Description:
	* @version:
	* @Author: Nan Kang
	* @Date: 2021-11-05 23:02:27
	* @LastEditTime: 2021-11-05 23:02:27
	* @LastEditors: Do not edit
	*/
   package main
   func main(){
	   str3:=
   }
   `
	fmt.Println(str3)
	// var n1 int32 = 12
	var n2 int64
	var n3 int8
	// n2 = n1 + 20
	// n3 = n1 + 20
	fmt.Println("n2=", n2, " n3=", n3)
}
