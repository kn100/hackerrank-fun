package main

import (
	"fmt"
)

func main() {
	testArr := []int32{0, 0, 1, 0, 0, 1, 0}
	fmt.Println(jumpingOnClouds(testArr))
}

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	finalCloud := len(c) - 1
	currCloud := 0
	numHops := 0
	for true {
		if (currCloud+2 <= len(c)-1) && c[currCloud+2] != 1 {
			currCloud = currCloud + 2
		} else {
			currCloud = currCloud + 1
		}
		numHops++
		if currCloud == finalCloud {
			break
		}
	}
	return int32(numHops)
}
