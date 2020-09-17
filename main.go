package main

import (
	"algo-6/sorting"
	"algo-6/utils"
	"fmt"
	"math/rand"
)

var testArr = []int{7, 0, 6, 1, 3, 2, 8, 5, 4, 9}

func main() {
	utils.GenerateFile("test", 100, 1, 100)
	fmt.Println("test", utils.BytesToInt16(utils.Read("test")))
	sorting.External{}.Sort("test", 8)

	//utils.SplitFile("test", "1", "2", 10)
	//fmt.Println(utils.BytesToInt16(utils.Read("2")))

	//sorting.MergeSort{}.Sort(testArr)
	//fmt.Println(testArr)
}

func randomString(l int) []int {
	bytes := make([]int, l)
	for i := 0; i < l; i++ {
		bytes[i] = randInt(65, 90)
	}
	return bytes
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
