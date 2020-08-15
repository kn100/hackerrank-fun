package main

import "fmt"

func main() {
	teststr := "UDDDUDUU"
	fmt.Println(countingValleys(8, teststr))
}

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	valleys := 0
	prevLevel := 0
	currLevel := 0
	for i, d := range s {
		if i != 0 {
			prevLevel = currLevel
		}
		if d == 'D' {
			currLevel--
		} else {
			currLevel++
		}
		if currLevel == 0 && prevLevel == -1 {
			valleys++
		}
	}
	return int32(valleys)
}
