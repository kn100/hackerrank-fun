package main

import "fmt"

func main() {
	fmt.Println(repeatedString("aa", 1))
}

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	if s == "" {
		return 0
	}
	stringLength := int64(len(s))
	numRepeats := int64(n / stringLength)
	remainingChars := int64(n % stringLength)

	c := countOccurrances(s)
	cp := countOccurrances(s[:remainingChars])
	return c*numRepeats + cp
}

func countOccurrances(s string) int64 {
	count := int64(0)
	for _, el := range s {
		if el == 'a' {
			count++
		}
	}
	return count
}
