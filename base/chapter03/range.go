package main

import "fmt"

func main() {

	// 这里我们使用range来计算一个切片的所有元素和
	// 这种方法对数组也适用
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println("index:", i)
	}
	fmt.Println("sum:", sum)

	// 使用range来遍历字典的时候，返回键值对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("%s -> %s\n", k, v)
		fmt.Println("sd")

	}

}
