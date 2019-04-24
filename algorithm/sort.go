package main

import (
	"cidychou/enjoy-go/algorithm/utlis"
	"fmt"
)

// 算法练习题   排序
func main() {
	numArr := []int{4, 2, 9, 45, 8, 36, 5, 1, 82}

	fmt.Println("排序前:", numArr)
	// numArr = bubbleSort(numArr)
	numArr = insertSort(numArr)
	fmt.Println("排序后:", numArr)

	//utlis.Fibonacci(10)
	fmt.Println("青蛙跳到4级有:", utlis.Fibonacci(50), "种跳法")

	fmt.Println("test")

}

// 1, 冒泡排序
func bubbleSort(numArr []int) []int {
	numCount := len(numArr)
	for i := 0; i < numCount; i++ {
		for j := numCount - 1; j > 0; j-- {
			if numArr[j] < numArr[j-1] {
				numArr[j], numArr[j-1] = numArr[j-1], numArr[j]
			}
		}
	}
	return numArr
}

// 2.快速排序
func quickSort(numArr []int) []int {
	return numArr
}

// 3.插入排序
func insertSort(numArr []int) []int {
	numCount := len(numArr)
	for i := 0; i < numCount; i++ {
		for j := i + 1; j < numCount; j++ {
			if numArr[i] > numArr[j] {
				numArr[i], numArr[j] = numArr[j], numArr[i]
			}
		}
	}
	return numArr
}
