package main

func Intersect(a , b []int) []int {
	m := make(map[int]uint8)
	for _, k := range a {
		m[k] |= (1 << 0)
	}
	for _, k := range b {
		m[k] |= (1 << 1)
	}
	var result []int
	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0
		if a && b {
			result = append(result, k)
		}
	}
	return result
}

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
