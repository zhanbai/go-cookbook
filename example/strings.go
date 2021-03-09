package main

import (
	"fmt"
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

func main() {
	stringIndex()
	substr()
}
