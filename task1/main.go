package main

import (
	"fmt"
	"strconv"
)

func main() {

	//只出现一次的数字
	arr := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6}
	result := func1(arr)
	fmt.Println("只出现一次的数字:", result)

	//判断是否是回文数
	a := 121
	b := 123
	fmt.Printf("%v是不是回文数%v \n", a, func2(a))
	fmt.Printf("%v是不是回文数%v \n", b, func2(b))

	//判断是否是有效符号
	strtrue := "(){}[]"
	strfalse := "{[]}("
	fmt.Printf("%v是不是有效符号%v \n", strtrue, func3(strtrue))
	fmt.Printf("%v是不是有效符号%v \n", strfalse, func3(strfalse))

	//查找公共前缀
	strs := []string{"flower", "flow", "flight"}
	strs2 := []string{"cat", "dog", "car"}

	fmt.Printf("strs前缀%v \n", func4(strs))
	fmt.Printf("strs2前缀%v \n", func4(strs2))

}

// 只出现一次的数字
func func1(arr []int) int {
	var result int = 0
	for _, a := range arr {
		result ^= a
	}
	return result
}

// 判断一个整数是否是回文数
func func2(a int) bool {
	if a < 0 {
		return false
	}
	//转换为字符串
	b := strconv.FormatInt(int64(a), 10)
	//字符数组倒转，转为新的字符串
	runes := []rune(b)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	str := string(runes)

	if str == b {
		return true
	} else {
		return false
	}
}

// 有效的括号
func func3(str string) bool {
	var i int = 0
	var j int = 0
	arr1 := "([{"
	arr2 := ")]}"
	//直接遍历字符串
	for _, v := range str {
		for _, r := range arr1 {
			if r == v {
				i++
			}
		}
		for _, r := range arr2 {
			if r == v {
				j++
			}
		}
	}

	return i == j

}

// 查找公共前缀
func func4(strs []string) string {
	if len(strs) == 0 {
		return "" // 如果字符串数组为空，直接返回空字符串
	}

	// 假设第一个字符串是最长公共前缀
	prefix := strs[0]

	// 遍历剩余的字符串
	for i := 1; i < len(strs); i++ {
		// 逐步缩短 prefix，直到它成为当前字符串的前缀
		for j := 0; j < len(prefix); j++ {
			// 如果 prefix 的当前字符与当前字符串的对应字符不匹配，
			// 或者当前字符串已经遍历完，则缩短 prefix
			if j >= len(strs[i]) || prefix[j] != strs[i][j] {
				prefix = prefix[:j] // 截取 prefix 的前 j 个字符
				break
			}
		}

		// 如果 prefix 已经为空，说明没有公共前缀，直接返回
		if prefix == "" {
			return ""
		}
	}
	return prefix
}
