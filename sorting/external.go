package sorting

import (
	"algo-6/utils"
	"fmt"
)

const int16Size = 2
const defaultTmp = "tmp"
const defaultMemoryUse = 1024

type External struct {
	nameCounter  *int
	TmpPath      string
	MaxMemoryUse int64
	InternalSort ISort
}

func (e External) Sort(path string) {
	counter := 0
	e.nameCounter = &counter

	if e.TmpPath == "" {
		e.TmpPath = defaultTmp
	}

	if e.MaxMemoryUse == 0 {
		e.MaxMemoryUse = defaultMemoryUse
	}

	if e.InternalSort == nil {
		e.sort(path)
	} else {
		e.sortWithInternal(path)
	}
}

func (e External) sortWithInternal(path string) {
	fileSize := utils.FileStatFromPath(path).Size()
	mem := e.MaxMemoryUse / int16Size

	if fileSize <= mem {
		targetFile := utils.OpenFile(path)
		buf := int64(0)

		if fileSize == mem {
			buf = mem
		} else {
			buf = fileSize
		}
		_, arr := utils.ReadAndParse(targetFile, buf)
		e.InternalSort.Sort(arr)

		utils.Seek(targetFile, 0)
		utils.Write(targetFile, utils.Int16ToBytes(arr))
		utils.CloseFile(targetFile)
		return
	}

	pathA := e.makeFileName()
	pathB := e.makeFileName()

	utils.SplitFile(path, pathA, pathB, e.MaxMemoryUse)

	e.sortWithInternal(pathA)
	e.sortWithInternal(pathB)

	e.mergeFiles(pathA, pathB, path)
}

func (e External) sort(path string) {
	fileSize := utils.FileStatFromPath(path).Size()

	if fileSize <= int16Size {
		return
	}

	pathA := e.makeFileName()
	pathB := e.makeFileName()

	utils.SplitFile(path, pathA, pathB, e.MaxMemoryUse)

	e.sort(pathA)
	e.sort(pathB)

	e.mergeFiles(pathA, pathB, path)
}

func (e External) mergeFiles(pathA, pathB, targetPath string) {
	fileA := utils.OpenFile(pathA)
	fileB := utils.OpenFile(pathB)
	targetFile := utils.OpenFile(targetPath)

	sizeA := utils.FileStat(fileA).Size()
	sizeB := utils.FileStat(fileB).Size()

	readA := int64(0)
	readB := int64(0)

	if sizeA > 0 && sizeB > 0 {
		_, a := utils.ReadAndParse(fileA, int64(int16Size))
		_, b := utils.ReadAndParse(fileB, int64(int16Size))

		for {
			if a[0] < b[0] {
				utils.Write(targetFile, utils.Int16ToBytes(a))
				readA += int16Size
				if readA == sizeA {
					utils.Write(targetFile, utils.Int16ToBytes(b))
					readB += int16Size
					break
				}
				_, a = utils.ReadAndParse(fileA, int64(int16Size))

			} else {
				utils.Write(targetFile, utils.Int16ToBytes(b))
				readB += int16Size
				if readB == sizeB {
					utils.Write(targetFile, utils.Int16ToBytes(a))
					readA += int16Size
					break
				}
				_, b = utils.ReadAndParse(fileB, int64(int16Size))

			}
		}
	}

	if readB < sizeB {
		maxBuf := e.maxReadBuf(readB, sizeB, e.MaxMemoryUse)
		read, b := utils.ReadAndParse(fileB, maxBuf)
		readB += read
		utils.Write(targetFile, utils.Int16ToBytes(b))
	}

	if readA < sizeA {
		maxBuf := e.maxReadBuf(readA, sizeA, e.MaxMemoryUse)
		read, a := utils.ReadAndParse(fileA, maxBuf)
		readA += read
		utils.Write(targetFile, utils.Int16ToBytes(a))
	}

	utils.CloseFile(targetFile)
	utils.CloseFile(fileA)
	utils.CloseFile(fileB)
	e.Remove([]string{pathA, pathB})
}

func (e External) maxReadBuf(read int64, fileSize int64, maxMemoryUse int64) (maxBuf int64) {
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

func (e External) makeFileName() string {
	*e.nameCounter++
	return fmt.Sprint(e.TmpPath, "/", *e.nameCounter)
}

func (e External) makeTmpName() string {
	return fmt.Sprint(e.TmpPath, "/", defaultTmp)
}

func (e External) Remove(files []string) {
	for _, v := range files {
		utils.Delete(v)
	}
}
