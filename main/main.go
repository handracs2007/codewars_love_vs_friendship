package main

import "fmt"

func WordsToMarks(s string) int {
	marks := 0

	for _, ch := range s {
		marks = marks + int(ch-'a'+1)
	}

	return marks
}

func main() {
	fmt.Println(WordsToMarks("love"))
	fmt.Println(WordsToMarks("friendship"))
}
