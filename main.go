package main

import (
	"algo-6/sorting"
	"algo-6/utils"
	"fmt"
	"math/rand"
	"time"
)

var testArr = []int{7, 0, 6, 1, 3, 2, 8, 5, 4, 9}

func main() {
	utils.GenerateFile("test", 200_000, 1, 10000)
	t := utils.BytesToInt16(utils.Read("test"))
	fmt.Println("test", utils.BytesToInt16(utils.Read("test")))

	x := make([]int, 0)

	for _, v := range t {
		x = append(x, int(v))
	}

	fmt.Println(sorting.MergeSort{}.Sort(x))
	s := time.Now()
	sorting.External{}.Sort("test", 1024)
	fmt.Println("test", utils.BytesToInt16(utils.Read("test")))
	fmt.Println("execution time", time.Since(s))
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
