package slice

func Intersection[K comparable](s1 []K, s2 []K) []K {
	var intersection []K
	hash := make(map[K]bool)
	for _, item := range s1 {
		hash[item] = true
	}
	for _, item := range s2 {
		if hash[item] {
			intersection = append(intersection, item)
		}
	}

	return intersection
}

func Reverse[K any](stack []K) []K {
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	return stack
}

func PeekFirst[K any](slice []K) K {
	return slice[0]
}

func Peek[K any](slice []K) K {
	return slice[len(slice)-1]
}

func Shift[K any](slice []K) ([]K, K) {
	item := slice[0]
	return slice[1:], item
}

func ShiftMany[K any](slice []K, amountToShift int) ([]K, []K) {
	items := slice[0:amountToShift]
	return slice[amountToShift:], items
}

func Prepend[K any](slice []K, item K) []K {
	slice = append(slice, *new(K))
	copy(slice[1:], slice)
	slice[0] = item
	return slice
}

func PrependMany[K any](slice []K, items []K) []K {
	newSlice := make([]K, len(items)+len(slice))
	copy(newSlice[:len(items)], items)
	copy(newSlice[len(items):], slice)
	return newSlice
}
