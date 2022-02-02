package main

import "fmt"

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, finaVal int) {
	//判断leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//先找中间下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > finaVal {
		//说明我们要查找的数,应该在 leftIndex --- middle-1
		BinaryFind(arr, leftIndex, middle-1, finaVal)
	} else if (*arr)[middle] < finaVal {
		BinaryFind(arr, middle+1, rightIndex, finaVal)
	} else {
		//找到了
		fmt.Printf("找到了值%v, 下标为%v \n", finaVal, middle)
	}

}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	//测试
	BinaryFind(&arr, 0, len(arr)-1, 9)
}
