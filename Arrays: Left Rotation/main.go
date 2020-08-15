package main

import "fmt"

func main() {
	testData := []int32{1, 2, 3, 4, 5}
	numRots := int32(4)
	fmt.Println(rotLeft(testData, numRots))
}

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	for i := int32(0); i < d; i++ {
		fv, tr := a[0], a[1:]
		a = append(tr, fv)
	}
	return a
}
