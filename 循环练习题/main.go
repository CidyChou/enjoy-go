package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	reverseSeed := []uint8{1, 1, 1, 1, 1, 12, 1, 1, 2, 1}
	fmt.Println(reverseSeed[8])

	//fmt.Println(twoSum([]int{3, 2, 4}, 6))
	//fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))

	//fmt.Println("checkPossibility", checkPossibility([]int{2, 3, 3, 2, 4}))
	//fmt.Println("climbStairs", climbStairs(3))
	// chapter1()
	// chapter2()
	// chapter3()
	// chapter4()
	// chapter5()
	// chapter6()
	// chapter7()
	// chapter8()

	//fmt.Println(threeSum([]int{-1, 0, 1, 2}))
}

// func threeSum(nums []int) [][]int {
// 	var res = make([][]int, 10000)
// 	index := 0
// 	for i := 0; i < len(nums); i++ {
// 		if i < len(nums)-2 {
// 			if nums[i]+nums[i+1]+nums[i+2] == 0 {
// 				res[index] = []int{nums[i], nums[i+1], nums[i+2]}
// 				index++
// 			}
// 		}
// 	}
// 	return res
// }

func threeSum(nums []int) [][]int {
	var res = make([][]int, 10)
	index := 0
	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j <= len(nums)-i; j++ {
			for h := j + 1; h <= len(nums)-j; h++ {
				fmt.Println(i, j, h)
			}
		}
		if i < len(nums)-2 {
			if nums[i]+nums[i+1]+nums[i+2] == 0 {
				res[index] = []int{nums[i], nums[i+1], nums[i+2]}
				index++
			}
		}
	}
	return res
}

func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	res := []int{0, 1, 2}

	for i := 3; i <= n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res[n]
}

/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */
func checkPossibility(nums []int) bool {
	var index int
	var position []int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			index++
			position = append(position, i)
		}
	}

	if index >= 2 {
		return false
	} else if index == 1 {
		p := position[0]
		if p != 0 && p != len(nums) && (nums[p+1] != nums[p] || nums[p-1] != nums[p]) {
			if nums[p+1]-nums[p-1] < 1 {
				return false
			}
		}
	}
	return true
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func uniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	var uniqueMorse []string
	var s string

	for _, word := range words {
		s = ""
		for _, l := range []rune(word) {
			s = s + morse[int(l)-97]
		}
		uniqueMorse = append(uniqueMorse, s)
	}

	return len(RemoveRepByLoop(uniqueMorse))
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop(slc []string) []string {
	result := []string{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 交换两个数
func exchangeString(a, b string) (string, string) {
	a, b = b, a
	return a, b
}

// 交换两个数
func chapter1() {
	a := 3
	b := 2
	a, b = b, a
	fmt.Println(a, b)
}

//给定3个整数，从小到大输出。
func chapter2() {
	a := 1
	b := 5
	c := 2

	if a > b {
		a, b = b, a
	}

	if a > c {
		a, c = c, a
	}

	if b > c {
		b, c = c, b
	}

	fmt.Println(a, b, c)

}

//打印9*9的乘法表
func chapter3() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

/*
水仙花数：三位数：[100,999]
每个位上的数字的立方和，刚好等于该数字本身。那么就叫水仙花数
比如：153
​111 + 555 + 333 = 1+125+27=153
*/
func chapter4() {
	for i := 100; i < 1000; i++ {
		a := i / 100
		b := i / 10 % 10
		c := i % 10
		if i == a*a*a+b*b*b+c*c*c {
			fmt.Println(i)
		}
	}
}

// 百元百鸡：一百元钱，买100只鸡。公鸡：每只5元，母鸡：每只3元，小鸡：1元3个。有多少种买法？
func chapter5() {
	/*
			思考：
		    公鸡：[0,20] a
		    母鸡：[0,33] b

			小鸡：100-a-b

	*/

	for i := 0; i <= 20; i++ {
		for j := 0; j <= 33; j++ {
			c := 100 - i - j
			if i*5+j*3+c == 100 && c%3 == 0 {
				fmt.Println(i, j, c)
			}
		}
	}

}

//  打印出2-100之间的素数
func chapter6() {
	/*
	   素数：也叫质数，只能被1和本身整除的数。
	   打印出2-100之间的质数。
	       7:
	           7%2,7%3,7%4,7%5,7%6,
	       8:
	           2,3,4,5,6,7

	*/

	for i := 2; i <= 100; i++ {
		flag := true //标记i是否是素数
		//循环验证是否是素数
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false //不是素数
				break
			}
		}
		//打印结果
		if flag == true {
			fmt.Println(i, "是素数")
		}
	}
}

//猜数游戏
//先由系统产生一个随机数，介于1-100之间。用户键盘输入正数来猜测数字，比较用户输入的整数和产生的随机数，给与用于提示：猜大了还是猜小了，用于继续猜测，直到猜对为止。
func chapter7() {
	/*
	   猜数游戏：
	   step1：由程序中产生一个随机数：[1,100]

	   step2：让用户键盘输入，猜测数字

	   step3：用户的猜测数字和随机数比较大小，给提示信息


	   随机数：38
	       70，提示大了
	       20，提示小了
	       50，提示大了
	       40，提示大了
	       35，提示小了
	       38，对啦。。
	*/
	//1.系统产生一个随机数
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1

	//用于输入数据
	a := 0
	for i := 1; ; i++ {
		//1.读取键盘输入
		fmt.Printf("请输入第 %d 次的猜测的数字：", i)
		fmt.Scanln(&a)
		//2.判断数值
		if a > num {
			fmt.Println("猜大了，下次小点。。")
		} else if a < num {
			fmt.Println("猜小了，下次大点。。。")
		} else {
			//相等
			fmt.Println("鼓掌，撒花。。。猜中了。。")
			break
		}

	}
}

func chapter8() {
	/*            行        空格    *        换行
	    *            1        4        1        1
	   ***            2        3        3        1
	  *****            3        2        5        1
	 *******        4        1        7        1
	*********        5        0        9        1
	 *******        1    2    1        7        1
	  *****            2    4    2        5        1
	   ***            3    6    3        3        1
	    *            4    8    4        1        1
	*/
	var n int
	fmt.Println("请输入菱形的边长：")
	fmt.Scanln(&n)
	//上三角
	for i := 1; i <= n; i++ { //控制行数
		//1.空格
		/*
		   i:1 j:=0,1,2,3        j<4
		   i:2,j:=0,1,2        j<3
		   i:3,j:=0,1            j<2
		*/
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		//2.*
		/*
		   i:1,j:0,            j<1    i:2
		   i:2,j:0,1,2            j<3    i:4
		   i:3,j:0,1,2,3,4        j<5    i:6
		   i:4,j:0,1,2,3,4,5,6    j<7    i:8
		*/
		for k := 0; k < i*2-1; k++ {
			fmt.Print("*")
		}
		//3.换行
		fmt.Println()

	}
	//下三角
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*n-1-i*2; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
