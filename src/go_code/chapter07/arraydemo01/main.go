package main

import (
	"fmt"
)

func main() {
	/**
	*有6只鸡,分别是3 5 1 3.4 2 50
	 */
	//思路分析:定义六个变量,分别表示6只鸡
	var hen1 float32 = 3
	var hen2 float32 = 5
	var hen3 float32 = 1
	var hen4 float32 = 3.4
	var hen5 float32 = 2
	var hen6 float32 = 50
	totalWeight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6
	// avgWeight := totalWeight / 6
	avgWeight := fmt.Sprintf("%.2f", totalWeight/6)

	fmt.Printf("totalWeight=%v avgWeight=%v\n", totalWeight, avgWeight)
	var hens [6]float64
	hens[0] = 3
	hens[1] = 5
	hens[2] = 1
	hens[3] = 3.4
	hens[4] = 2
	hens[5] = 50
	totalWeight2 := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight2 += hens[i]
	}
	avgWeight2 := fmt.Sprintf("%.2f", totalWeight2/float64(len(hens)))
	fmt.Printf("totalWeight2=%v avgWeight2=%v\n", totalWeight2, avgWeight2)

}
