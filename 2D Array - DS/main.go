package main

import (
	"fmt"
	"math"
)

func main() {
	testData := [][]int32{
		{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, 1, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 0, 2, 4, 0},
	}
	fmt.Println(hourglassSum(testData))
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	max := math.MinInt32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			hv := calcHourglass(i, j, arr)
			if max < hv {
				max = hv
			}
		}
	}
	return int32(max)
}

func calcHourglass(i int, j int, arr [][]int32) int {
	return int(arr[i+0][j+0] + arr[i+0][j+1] + arr[i+0][j+2] + arr[i+1][j+1] + arr[i+2][j+0] + arr[i+2][j+1] + arr[i+2][j+2])
}
