package sorting

const maxRank = int16(5)

type Radix struct {
}

func (r Radix) Sort(arr []int16) []int16 {

	tmp := make([]int16, len(arr))
	rank := int16(1)

	for rank <= maxRank {
		ranks := make([]int64, 10)

		for _, v := range arr {
			d := r.getDigit(v, rank)
			ranks[d]++
		}

		for i := 1; i < len(ranks); i++ {
			ranks[i] = ranks[i] + ranks[i-1]
		}

		for i := len(arr) - 1; i >= 0; i-- {
			d := r.getDigit(arr[i], rank)
			ranks[d]--
			tmp[ranks[d]] = arr[i]
		}

		copy(arr, tmp)
		rank++
	}

	return arr
}

func (r Radix) getDigit(num int16, rank int16) int16 {
	digit := int16(0)

	for rank > 0 {
		digit = num % 10
		num /= 10
		rank--
	}

	return digit
}
