package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 1.1 访问子串
func stringIndex() {
	index := strings.Index("bai.zhan@qq.com", "@")
	if index == -1 {
		fmt.Println("There was no @ in the e-mail address!")
	}
}

// 1.2 抽取子串
func substr() {
	fmt.Println("zhanbai"[0:4])
	fmt.Println("watch out for that tree"[6:11])
	fmt.Println("watch out for that tree"[17:])
}

// 1.3 替换子串
func substrReplace() {
	str := "My pet is a blue dog."
	fmt.Println(strings.Replace(str, str[12:], "fish.", 1))
	fmt.Println(strings.Replace(str, str[12:16], "green", 1))
	creditCard := "4111 1111 1111 1111"
	fmt.Println(strings.Replace(creditCard, creditCard[0:len(creditCard)-4], "xxxx ", 1))
}

// 1.4 逐字节处理字符串
func forStr() {
	str := "This weekend, I'm going shopping for a pet chicken."
	vowels := 0
	for i := 0; i < len(str); i++ {
		if strstr("aeiouAEIOU", str[i:i+1]) != "" {
			vowels++
		}
	}
	fmt.Println(vowels)
}

func strstr(haystack string, needle string) string {
	if needle == "" {
		return ""
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return ""
	}
	return haystack[idx+len([]byte(needle))-1:]
}

// Look and Say 序列
func forLookandsay()  {
	s := strconv.Itoa(1)
	fmt.Println(s)

	for i := 0; i < 10; i++ {
		s = lookandsay(s)
		fmt.Println(s)
	}
}

func lookandsay(s string) string {
	// 将返回值初始化为一个空串
	r := ""
	// m 包含要统计的字符，初始化为字符串中的第一个字符
	m := s[0:1]
	
	// n 是已经查看过的 m 个数，初始化为 1
	n := 1
	for i := 1; i < len(s); i++ {
		// 如果这个字符与上一个相同
		if s[i:i+1] == m {
			// 将这个字符的个数增 1
			n++
		} else {
			// 否则，将字符个数和字符本身增加到返回值
			r += strconv.Itoa(n) + m
			// 将要查找的字符设置为当前字符
			m = s[i:i+1]
			// 并将字符个数重置为 1
			n = 1
		}
	}
	// 返回构建的字符串以及最后的字符个数及字符
	return r + strconv.Itoa(n) + m
}

// 1.5 按单词或字节反转字符串
func revStr()  {
	fmt.Println(strrev("This is not a palindrome."))
}

func strrev(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}	
	return string(runes[n:])
}

func revArr()  {
	s := "Once upon a time there was a turtle."
	// 将字符串分解为单词
	words := strings.Split(s, " ")
	// 反转单词数组
	words = arrayReverse(words)
	s = strings.Join(words, " ")
	fmt.Println(s)
}

// 反转单词数组
func arrayReverse(s []string) []string {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	} 
	return s
}

func main() {
	// stringIndex()
	// substr()
	// substrReplace()
	// forStr()
	// forLookandsay()
	// revStr()
	revArr()
}
