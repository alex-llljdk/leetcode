package main

import (
	"fmt"
	"strconv"
)

/*  模拟法，检测每个位置的长度进行相加，通过求余得出该位置的数值，通过求取累加值/2得出进位值 */
// func addBinary(a string, b string) string {
// 	ans := ""
// 	len1 := len(a)
// 	len2 := len(b)
// 	carry := 0
// 	n := max(len1, len2)

// 	for i := 0; i < n; i++ {
// 		if i < len1 {
// 			carry += int(a[len1-i-1] - '0')
// 		}
// 		if i < len2 {
// 			carry += int(b[len2-i-1] - '0')
// 		}
// 		ans = strconv.Itoa(carry%2) + ans
// 		carry /= 2
// 	}
// 	if carry > 0 {
// 		ans = "1" + ans
// 	}
// 	return ans
// }

/* 先转为int类型相加，再转为2进制代码 */
func addBinary(a string, b string) string {
	var num1, num2 int64 = 0, 0
	var err error
	num1, err = strconv.ParseInt(a, 2, 64)
	num2, err = strconv.ParseInt(b, 2, 64)
	fmt.Println(num1, num2)
	if err != nil {
		return ""
	}
	res := num1 + num2
	fmt.Println(res)
	return strconv.FormatInt(int64(res), 2)
}

func main() {
	fmt.Println(addBinary("101", "1"))
}
