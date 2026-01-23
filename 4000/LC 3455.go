func shortestMatchingSubstring(s string, p string) int {
    t := strings.Split(p, "*")
    a, b, c := t[0], t[1], t[2]
    pos1 := kmp(s, a)
    pos2 := kmp(s, b)
    pos3 := kmp(s, c)

    ans := math.MaxInt32 
    i, k := 0, 0
    for _, j := range pos2 {
        for k < len(pos3) && pos3[k] < len(b) + j {
            k++
        }
        if k == len(pos3) {
            break
        }

        for i < len(pos1) && pos1[i] <= j - len(a) {
            i++
        }
        if i > 0 {
            ans = min(ans, pos3[k] + len(c) - pos1[i - 1])
        }
    } 

    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}

func kmp(text, pattern string) (pos []int) {
	m := len(pattern)
    if m == 0 {
        pos = make([]int, len(text) + 1)
        for i := range pos {
            pos[i] = i
        } 
        return
    }

	pi := make([]int, m)
	cnt := 0
	for i := 1; i < m; i++ {
		v := pattern[i]
		for cnt > 0 && pattern[cnt] != v {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == v {
			cnt++ 
		}
		pi[i] = cnt
	}

	cnt = 0
	for i, v := range text {
		for cnt > 0 && pattern[cnt] != byte(v) {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == byte(v) {
			cnt++
		}
		if cnt == m {
			pos = append(pos, i-m+1)
			cnt = pi[cnt-1]
		}
	}
	return
}