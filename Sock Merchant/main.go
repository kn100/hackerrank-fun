package main

import "fmt"

func main() {
	numSocks := 9
	pairs := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(int32(numSocks), pairs))

}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	countedSocks := countSocks(ar)
	return countPairs(countedSocks)
}

func countSocks(ar []int32) map[int32]int32 {
	counts := make(map[int32]int32)
	for _, sock := range ar {
		if _, ok := counts[sock]; ok {
			// we've seen this sock before
			counts[sock]++
		} else {
			// Brand new sock!
			counts[sock] = 1
		}
	}
	return counts
}

func countPairs(ma map[int32]int32) int32 {
	pairCount := 0
	for _, count := range ma {
		pairCount = pairCount + (int(count) / 2)
	}
	return int32(pairCount)
}
