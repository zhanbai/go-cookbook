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

func main() {
	stringIndex()
}
