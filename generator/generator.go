package generator

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"unsafe"
)

const min = 0
const max = 65535

func GenerateFile(path string, amount int) error {
	f, err := os.Create(path)
	defer closeFile(f)

	if err != nil {
		return err
	}

	arr := make([]byte, 0, amount*2)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := int64(0); i < int64(amount); i++ {

		num := min + rand.Intn(max-min)

		fmt.Println(num)

		b := *(*byte)(unsafe.Pointer(&num))
		arr = append(arr, b)
		b = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(1)))
		arr = append(arr, b)
	}

	_, err = f.Write(arr)

	buf, err := Read(path)

	fmt.Println(Parse(buf))
	return nil
}

func Read(path string) ([]byte, error) {
	f, err := os.Open(path)

	defer closeFile(f)

	buf := make([]byte, 0)

	if err != nil {
		return buf, err
	}

	s, err := f.Stat()

	if err != nil {
		return buf, err
	}

	buf = make([]byte, s.Size())

	_, err = f.Read(buf)

	return buf, err
}

func Parse(buf []byte) []int {
	res := make([]int, 0, len(buf)/2)

	for i := 0; i < len(buf); i++ {
		n := 0
		*(*byte)(unsafe.Pointer(&n)) = buf[i]
		i++
		*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&n)) + uintptr(1))) = buf[i]

		res = append(res, n)
	}

	return res
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
