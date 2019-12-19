package main

import "fmt"

func main() {
	fmt.Println(BruteForce("helloworldww", "world"))
	fmt.Println(BruteForce("helloworlhww", "world"))
}

func BruteForce(str string, pattern string) int {
	strLength := len(str)
	patternLength := len(pattern)

	for i := 0; i < strLength-patternLength; i++ {
		if str[i:i+patternLength] == pattern {
			return i
		}
	}
	return -1
}
