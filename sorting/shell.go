package sorting

import (
	"math"
)

const KnutSteps = "knutSteps"
const SedgewickSteps = "sedgewickSteps"

type Shell struct {
	StepType string
}

func (s Shell) Sort(a []int16) []int16 {
	switch s.StepType {
	case KnutSteps:
		return s.sortWithSteps(a, s.knutSteps(len(a)))
	case SedgewickSteps:
		return s.sortWithSteps(a, s.sedgewickSteps(len(a)))
	default:
		return s.sortWithShellSteps(a)
	}
}

func (s Shell) sortWithShellSteps(a []int16) []int16 {
	for step := len(a) / 2; step > 0; step /= 2 {
		for i := step; i < len(a); i++ {
			for j := i - step; j >= 0 && a[j] > a[j+step]; j -= step {
				a[j], a[j+step] = a[j+step], a[j]
			}
		}
	}

	return a
}

func (s Shell) sortWithSteps(a []int16, steps []int) []int16 {
	for k := len(steps) - 1; k >= 0; k-- {
		step := steps[k]
		for i := step; i < len(a); i++ {
			for j := i - step; j >= 0 && a[j] > a[j+step]; j -= step {
				a[j], a[j+step] = a[j+step], a[j]
			}
		}
	}

	return a
}

func (s Shell) knutSteps(l int) []int {
	steps := make([]int, 0)
	p := float64(1)
	const c = float64(3)
	maxStep := l / 3
	step := int(0)

	for true {
		step = int((math.Pow(c, p) - 1) / 2)
		p++
		if step < maxStep {
			steps = append(steps, step)
		} else {
			break
		}
	}

	return steps
}

func (s Shell) sedgewickSteps(l int) []int {
	steps := []int{1}
	const four = float64(4)
	const three = float64(3)
	const two = float64(2)

	p := float64(1)
	step := 0

	for true {
		step = int(math.Pow(four, p) + (three * math.Pow(two, p-1)) + 1)
		p++
		if step < l {
			steps = append(steps, step)
		} else {
			break
		}
	}

	return steps
}
