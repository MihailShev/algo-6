package sorting

type MergeSort struct {
}

func (m MergeSort) Sort(arr []int16) []int16 {
	size := len(arr)

	if size <= 1 {
		return arr
	}

	leftSize := size / 2

	m.Sort(arr[:leftSize])
	m.Sort(arr[leftSize:])

	leftIndex := 0
	rightIndex := leftSize

	tmp := make([]int16, 0, size)

	for leftIndex < leftSize && rightIndex < size {
		if arr[leftIndex] < arr[rightIndex] {
			tmp = append(tmp, arr[leftIndex])
			leftIndex++
		} else {
			tmp = append(tmp, arr[rightIndex])
			rightIndex++
		}
	}

	if leftIndex == leftSize {
		tmp = append(tmp, arr[rightIndex:]...)
	}

	if rightIndex == size {
		tmp = append(tmp, arr[leftIndex:leftSize]...)
	}

	copy(arr, tmp)

	return arr
}
