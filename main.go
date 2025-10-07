package main

import (
	"bufio"
	"fmt"	
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	var someString string
	fmt.Fscan(reader, &someString)

	resultFlag := uniqueCharsInString(someString)
	fmt.Print(resultFlag)
}

func uniqueCharsInString(s string) bool {
	lowerS := strings.ToLower(s)	
	runeString := []rune(lowerS)
	chars := make(map[rune]int, len(runeString))

	resultFlag := true

	for _, char := range runeString {
		if _, ok := chars[char]; ok {
			resultFlag = false
			break
		}
		chars[char]++
	}

	return resultFlag
}
