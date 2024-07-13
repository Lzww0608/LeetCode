func bestSeqAtIndex(height []int, weight []int) int {
    n := len(height)
	pos := make([]int, n)
	for i := range pos {
		pos[i] = i
	}

	sort.Slice(pos, func(i, j int) bool {
		if height[pos[i]] == height[pos[j]] {
			return weight[pos[i]] > weight[pos[j]] 
		}
		return height[pos[i]] < height[pos[j]]
	})

	st := []int{}
	for _, x := range pos {
		t := weight[x]
		idx := sort.Search(len(st), func(i int) bool {
			return st[i] >= t
		})
		if idx < len(st) {
			st[idx] = t
		} else {
			st = append(st, t)
		}
	}

	return len(st)
}