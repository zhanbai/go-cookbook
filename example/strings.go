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

func main() {
	stringIndex()
	substr()
}
