package sorting

type Quick struct {
}

func (q Quick) Sort(arr []int16) []int16 {
	q.sort(arr, 0, len(arr)-1)
	return arr
}

func (q Quick) sort(arr []int16, l, r int) {
	if l >= r {
		return
	}

	a := q.partition(arr, l, r)
	q.sort(arr, l, a-1)
	q.sort(arr, a+1, r)
}

func (q Quick) partition(arr []int16, l, r int) int {
	pivot := arr[r]
	a := l - 1
	for m := l; m <= r; m++ {
		if arr[m] <= pivot {
			a++
			arr[a], arr[m] = arr[m], arr[a]
		}
	}

	return a
}
