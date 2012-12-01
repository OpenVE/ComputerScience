// Daniel R. (sadasant.com)
// 05/10/2012
//
// License:
//   GNU GPL 3.0
//

package sort

func Insertion(array []int) []int {
	var now int
	answer := make([]int, len(array))
	copy(answer, array)
	for i, v := range answer {
		now = i
		for now > 0 && answer[now-1] > v {
			answer[now] = answer[now-1]
			now--
		}
		answer[now] = v
	}
	return answer
}

func Bubble(array []int) []int {
	var current, prev int
	var swapped bool
	l := len(array)
repeat:
	swapped = false
	for i := 1; i < l; i++ {
		current = array[i]
		prev = array[i-1]
		if prev > current {
			array[i] = prev
			array[i-1] = current
			swapped = true
		}
	}
	if swapped {
		goto repeat
	}
	return array
}

func Quick(array []int) []int {
	if (len(array) <= 1) {
		return array
	}
	var less, equal, greater []int
	pivot := array[0]
	for _, v := range array[0:] {
		switch {
		case v < pivot:
			less = append(less, v)
		case v == pivot:
			equal = append(equal, v)
		case v > pivot:
			greater = append(greater, v)
		}
	}
	mix := make([]int, 0)
	mix = append(mix, Quick(less)...)
	mix = append(mix, equal...)
	return append(mix, Quick(greater)...)
}

func Merge(array []int) []int {
	l := len(array)
	if l < 2 {
		return array
	}
	half := l / 2
	return merge(Merge(array[:half]), Merge(array[half:]))
}
func merge(a, b []int) []int {
	la, lb := len(a), len(b)
	if la == 0 {
		return b
	}
	if lb == 0 {
		return a
	}
	mix := make([]int, 1)
	switch {
	case a[0] < b[0]:
		mix[0] = a[0]
		return append(mix, merge(a[1:], b)...)
	default:
		mix[0] = b[0]
		return append(mix, merge(b[1:], a)...)
	}
	return mix
}
