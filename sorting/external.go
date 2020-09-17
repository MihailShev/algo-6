package sorting

import (
	"algo-6/utils"
	"fmt"
)

const defaultTmp = "tmp"

type External struct {
	nameCounter *int
	TmpPath     string
}

func (e External) Sort(path string, maxMemoryUse int64) {
	counter := 0
	e.nameCounter = &counter

	if e.TmpPath == "" {
		e.TmpPath = defaultTmp
	}

	e.sort(path, maxMemoryUse)
}

func (e External) sort(path string, maxMemory int64) string {
	fileSize := utils.FileStatFromPath(path).Size()

	if fileSize <= 2 {
		return path
	}

	pathA := e.makeFilePath()
	pathB := e.makeFilePath()

	utils.SplitFile(path, pathA, pathB, maxMemory)

	fmt.Println("a", utils.BytesToInt16(utils.Read(pathA)))
	fmt.Println("b", utils.BytesToInt16(utils.Read(pathB)))

	pathA = e.sort(pathA, maxMemory)
	pathB = e.sort(pathB, maxMemory)

	mergePath := e.makeFilePath()

	fileA := utils.OpenFile(pathA)
	fileB := utils.OpenFile(pathB)
	mergeFile := utils.CreateFile(mergePath)
	defer utils.CloseFile(mergeFile)

	sizeA := utils.FileStat(fileA).Size()
	sizeB := utils.FileStat(fileB).Size()
	maxMemoryForRead := maxMemory / 2
	readA := int64(0)
	readB := int64(0)

	for readA < sizeA && readB < sizeB {
		read, a := utils.ReadAndParse(fileA, e.calcMaxBuf(readA, sizeA, maxMemoryForRead))
		readA += read

		read, b := utils.ReadAndParse(fileB, e.calcMaxBuf(readB, sizeB, maxMemoryForRead))
		readB += read

		utils.Write(mergeFile, utils.Int16ToBytes(e.merge(a, b)))
	}
	utils.CloseFile(fileA)
	utils.CloseFile(fileB)
	e.Remove([]string{pathA, pathB})

	fmt.Println("res", utils.BytesToInt16(utils.Read(mergePath)))

	return mergePath
}

func (e External) calcMaxBuf(read int64, fileSize int64, maxMemoryUse int64) (maxBuf int64) {
	if read+maxMemoryUse > fileSize {
		maxBuf = fileSize - read
	} else {
		maxBuf = maxMemoryUse
	}

	if maxBuf%2 != 0 {
		maxBuf--
	}

	return maxBuf
}

func (e External) makeFilePath() string {
	*e.nameCounter++
	return fmt.Sprint(e.TmpPath, "/", *e.nameCounter)
}

func (e External) getTmpName() string {
	return fmt.Sprint(e.TmpPath, "/", defaultTmp)
}

func (e External) merge(a []int16, b []int16) []int16 {
	aIndex := 0
	bIndex := 0
	aSize := len(a)
	bSize := len(b)
	res := make([]int16, 0, aSize+bSize)

	for aIndex < aSize && bIndex < bSize {
		if a[aIndex] < b[bIndex] {
			res = append(res, a[aIndex])
			aIndex++
		} else {
			res = append(res, b[bIndex])
			bIndex++
		}
	}

	if aIndex == aSize {
		res = append(res, b[bIndex:]...)
	}

	if bIndex == bSize {
		res = append(res, a[aIndex:]...)
	}

	return res
}

func (e External) Remove(files []string) {
	for _, v := range files {
		utils.Delete(v)
	}
}
