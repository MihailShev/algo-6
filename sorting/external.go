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

func (e External) sort(path string, maxMemory int64) {
	fileSize := utils.FileStatFromPath(path).Size()

	if fileSize <= 2 {
		return
	}

	pathA := e.makeFilePath()
	pathB := e.makeFilePath()

	utils.SplitFile(path, pathA, pathB, maxMemory)

	e.sort(pathA, maxMemory)
	e.sort(pathB, maxMemory)

	fileA := utils.OpenFile(pathA)
	fileB := utils.OpenFile(pathB)
	targetFile := utils.OpenFile(path)

	sizeA := utils.FileStat(fileA).Size()
	sizeB := utils.FileStat(fileB).Size()

	readA := int64(0)
	readB := int64(0)

	if sizeA > 0 && sizeB > 0 {
		_, a := utils.ReadAndParse(fileA, int64(2))
		_, b := utils.ReadAndParse(fileB, int64(2))

		for {
			if a[0] < b[0] {
				utils.Write(targetFile, utils.Int16ToBytes(a))
				readA += 2
				if readA == sizeA {
					utils.Write(targetFile, utils.Int16ToBytes(b))
					readB += 2
					break
				}
				_, a = utils.ReadAndParse(fileA, int64(2))

			} else {
				utils.Write(targetFile, utils.Int16ToBytes(b))
				readB += 2
				if readB == sizeB {
					utils.Write(targetFile, utils.Int16ToBytes(a))
					readA += 2
					break
				}
				_, b = utils.ReadAndParse(fileB, int64(2))

			}
		}
	}

	if readB < sizeB {
		maxBuf := e.calcMaxBuf(readB, sizeB, maxMemory)
		read, b := utils.ReadAndParse(fileB, maxBuf)
		readB += read
		utils.Write(targetFile, utils.Int16ToBytes(b))
	}

	if readA < sizeA {
		maxBuf := e.calcMaxBuf(readA, sizeA, maxMemory)
		read, a := utils.ReadAndParse(fileA, maxBuf)
		readA += read
		utils.Write(targetFile, utils.Int16ToBytes(a))
	}

	utils.CloseFile(targetFile)
	utils.CloseFile(fileA)
	utils.CloseFile(fileB)
	e.Remove([]string{pathA, pathB})
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
