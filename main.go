package main

import (
	"algo-6/sorting"
	"algo-6/utils"
	"fmt"
	"math"
	"time"
)

const tmpPath = "tmp/test"

func main() {
	src1e6 := "tmp/1e6"
	src1e7 := "tmp/1e7"
	src1e8 := "tmp/1e8"
	src1e9 := "tmp/1e9"

	fmt.Println("Start generating files")
	utils.GenerateFile(src1e6, 1_000_000, 1, math.MaxInt16)
	fmt.Println("1e6")
	utils.GenerateFile(src1e7, 10_000_000, 1, math.MaxInt16)
	fmt.Println("1e7")
	utils.GenerateFile(src1e8, 100_000_000, 1, math.MaxInt16)
	fmt.Println("1e8")
	utils.GenerateFile(src1e9, 1000_000_000, 1, math.MaxInt16)
	fmt.Println("1e9")

	ext1e6 := "tmp/external_1e6-test"
	extWithInternal1e6 := "tmp/ext_1e6-with-internal-sort-test"
	extWithInternal1e7 := "tmp/ext_1e7-with-internal-sort-test"
	extWithInternal1e8 := "tmp/ext_1e8-with-internal-sort-test"

	fmt.Println("copying...")

	utils.Copy(ext1e6, src1e6)
	utils.Copy(extWithInternal1e6, src1e6)
	utils.Copy(extWithInternal1e7, src1e7)
	utils.Copy(extWithInternal1e8, src1e8)
	fmt.Println("Finished")

	//fmt.Printf("\n\n*** Test external sort 1e6 ***\n\n")
	//test(func() {Л
	//	sorting.External{}.Sort(ext1e6)
	//})

	fmt.Printf("*** Test external sort with internal sort 1e6 ***\n\n")
	test(func() {
		sorting.External{MaxMemoryUse: 4096, InternalSort: sorting.Shell{
			StepType: sorting.SedgewickSteps,
		}}.Sort(extWithInternal1e6)
	})

	fmt.Printf("*** Test external sort with internal sort 1e7 ***\n\n")
	test(func() {
		sorting.External{MaxMemoryUse: 4096, InternalSort: sorting.Shell{
			StepType: sorting.SedgewickSteps,
		}}.Sort(extWithInternal1e7)
	})

	//fmt.Printf("*** Test external sort with internal sort 1e8 ***\n\n")
	//test(func() {
	//	sorting.External{MaxMemoryUse: 4096, InternalSort: sorting.Shell{
	//		StepType: sorting.SedgewickSteps,
	//	}}.Sort(extWithInternal1e8)
	//})

	fmt.Printf("*** Test radix sort 1e6 ***\n\n")
	num := utils.ReadPathAndParse(src1e6)
	test(func() {
		sorting.Radix{}.Sort(num)
	})

	fmt.Printf("*** Test radix sort 1e7 ***\n\n")
	num = utils.ReadPathAndParse(src1e7)
	test(func() {
		sorting.Radix{}.Sort(num)
	})

	fmt.Printf("*** Test radix sort 1e8 ***\n\n")
	num = utils.ReadPathAndParse(src1e8)
	test(func() {
		sorting.Radix{}.Sort(num)
	})

	fmt.Printf("*** Test radix sort 1e9 ***\n\n")
	num = utils.ReadPathAndParse(src1e9)
	test(func() {
		sorting.Radix{}.Sort(num)
	})
}

func test(run func()) {
	start := time.Now()
	run()
	stop := time.Since(start)
	fmt.Printf("Execution time: %s\n\n", stop)
}
