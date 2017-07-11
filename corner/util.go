package corner

func minIntSlice(slice []int) int {
	m := slice[0]
	for _, x := range slice {
		if x < m {
			m = x
		}
	}
	return m
}

func maxIntSlice(slice []int) int {
	m := slice[0]
	for _, x := range slice {
		if x > m {
			m = x
		}
	}
	return m
}
