package floatlist

import (
	"sort"
)

type Floatlist []float64

func (list Floatlist) Mean() float64 {
	length := float64(len(list))
	total := list.Sum()
	return total / length
}

func (list Floatlist) Sum() float64 {
	total := 0.0
	for _, f := range list {
		total += f
	}
	return total
}

func (list Floatlist) Median() float64 {
	// make sure it's sorted
	sort.Float64s(list)

	length := len(list)

	var index int
	if length%2 == 0 {
		index = (length / 2) - 1
	} else {
		index = (length - 1) / 2
	}

	return list[index]

}

func (list Floatlist) GetCountsByValue() map[float64]int {
	counts := map[float64]int{}

	for _, f := range list {
		if _, ok := counts[f]; !ok {
			counts[f] = 0
		}
		counts[f]++
	}

	return counts
}

func (list Floatlist) GetCountByValue(val float64) int {
	counts := list.GetCountsByValue()

	if count, ok := counts[val]; ok {
		return count
	}
	return 0
}

func (list Floatlist) Mode() float64 {

	counts := list.GetCountsByValue()
	var max float64
	maxCount := 0

	for val, count := range counts {
		if count > maxCount {
			max = val
			maxCount = count
		}
	}

	return max
}
