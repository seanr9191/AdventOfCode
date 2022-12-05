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
