package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// q3_1()
	// q3_2()
	// q3_3()
	q3_4()
}

func q3_1() {
	k := 0
I:
	for i := 0; i < 100; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("A")
			if k > 100 {
				break I
			}
			k++
		}
		fmt.Println("")
	}
}

func q3_2() {
	str := "asSASA ddd 啊带伞dsjkdsjs dk"
	count := utf8.RuneCount([]byte(str))
	strRune := []rune(str)
	bCount := 0

	for i := 0; i < count; i++ {
		bCount += utf8.RuneLen(strRune[i])
	}
	fmt.Println("字符串的字符数为：", count)
	fmt.Println("字符串的字节数为：", bCount)
}

func q3_3() {
	str := "asSASA ddd 啊带伞dsjkdsjs dk"
	strRune := []rune(str)
	strRep := []rune("abc")
	for i := 0; i < 3; i++ {
		strRune[3+i] = strRep[i]
	}
	strNew := string(strRune)
	fmt.Print(strNew)

}

func q3_4() {
	str := "foobar"
	strRune := []rune(str)
	strNew := ""
	for i := utf8.RuneCountInString(str) - 1; i >= 0; i-- {
		strNew += string(strRune[i])
	}
	fmt.Println(strNew)
}
