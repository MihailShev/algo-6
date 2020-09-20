package generator

import (
	"algo-6/utils"
	"math"
	"math/rand"
	"os"
	"time"
)

const min = 0
const max = math.MaxInt16

func GenerateFile(path string, amount int16) error {
	f, err := os.Create(path)
	defer utils.CloseFile(f)

	if err != nil {
		return err
	}

	arr := make([]int16, 0, amount)

	rand.Seed(time.Now().UTC().UnixNano())

	i := 0

	for i < int(amount) {
		num := int16(min + rand.Intn(max-min))
		arr = append(arr, num)
	}

	utils.Write(f, utils.Int16ToBytes(arr))

	return nil
}
