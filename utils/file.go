package utils

import (
	"io"
	"math/rand"
	"os"
	"time"
	"unsafe"
)

func GenerateFile(path string, amount int, min int, max int) {
	f, err := os.Create(path)
	defer CloseFile(f)

	handleError(err)

	arr := make([]byte, 0, amount*2)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := int64(0); i < int64(amount); i++ {

		num := min + rand.Intn(max-min)

		b := *(*byte)(unsafe.Pointer(&num))
		arr = append(arr, b)
		b = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(1)))
		arr = append(arr, b)
	}

	_, err = f.Write(arr)
	handleError(err)
}

func Write(f *os.File, buf []byte) {
	_, e := f.Write(buf)

	handleError(e)
}

func Delete(path string) {
	err := os.Remove(path)
	handleError(err)
}

func ReadPathAndParse(path string) []int16 {
	f := OpenFile(path)
	defer CloseFile(f)
	size := FileStat(f).Size()
	_, res := ReadAndParse(f, size)
	return res
}

func ReadBuf(f *os.File, maxBuf int64) (int64, []byte) {
	buf := make([]byte, maxBuf)

	r, err := f.Read(buf)
	handleError(err)
	return int64(r), buf
}

func ReadAndParse(f *os.File, n int64) (r int64, p []int16) {
	r, b := ReadBuf(f, n)
	p = BytesToInt16(b)
	return r, p
}

func Int16ToBytes(nums []int16) []byte {
	bytes := make([]byte, 0, len(nums)*2)
	for _, v := range nums {

		b := *(*byte)(unsafe.Pointer(&v))
		bytes = append(bytes, b)
		b = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&v)) + uintptr(1)))
		bytes = append(bytes, b)
	}

	return bytes
}

func BytesToInt16(buf []byte) []int16 {
	res := make([]int16, 0, len(buf)/2)

	for i := 0; i < len(buf); i++ {
		n := int16(0)
		*(*byte)(unsafe.Pointer(&n)) = buf[i]
		i++
		*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&n)) + uintptr(1))) = buf[i]

		res = append(res, n)
	}

	return res
}

func SplitFile(path string, pathA, pathB string, maxBuf int64) {
	source := OpenFile(path)
	defer CloseFile(source)
	fileA := CreateFile(pathA)
	defer CloseFile(fileA)
	fileB := CreateFile(pathB)
	defer CloseFile(fileB)
	sourceSize := FileStat(source).Size()

	aSize := sourceSize / 2

	if aSize%2 != 0 {
		aSize--
	}

	bSize := sourceSize - aSize

	limitCopy(fileA, source, aSize, maxBuf)
	limitCopy(fileB, source, bSize, maxBuf)
}

func Copy(dst, source string) {
	dstFile := CreateFile(dst)
	sourceFile := OpenFile(source)
	_, err := io.Copy(dstFile, sourceFile)
	handleError(err)
	CloseFile(dstFile)
	CloseFile(sourceFile)
}

func limitCopy(dst *os.File, source *os.File, size int64, maxBuf int64) {
	copied := int64(0)
	toCopy := int64(0)

	for {

		if copied+maxBuf > size {
			toCopy = size - copied
		} else {
			toCopy = maxBuf
		}

		w, err := io.CopyN(dst, source, toCopy)

		if err == io.EOF {
			break
		} else {
			handleError(err)
		}

		copied += w

		if copied >= size {
			break
		}
	}
}

func CreateFile(path string) *os.File {
	f, err := os.Create(path)
	handleError(err)
	return f
}

func OpenFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_RDWR, 0660)
	handleError(err)
	return f
}

func FileStatFromPath(path string) os.FileInfo {
	f, err := os.Open(path)
	defer CloseFile(f)

	handleError(err)

	return FileStat(f)
}

func FileStat(f *os.File) os.FileInfo {
	s, err := f.Stat()

	handleError(err)

	return s
}

func Seek(f *os.File, offset int64) {
	_, err := f.Seek(offset, 0)
	handleError(err)
}

func CloseFile(f *os.File) {
	err := f.Close()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
