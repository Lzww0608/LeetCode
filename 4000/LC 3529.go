func kmp(text []byte, pattern string) (pos []int) {
	m := len(pattern)
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

func countCells(grid [][]byte, pattern string) (ans int) {
    m, n := len(grid), len(grid[0])
    s := make([]byte, m * n)
    for i := range m {
        for j := range n {
            s[i * n + j] = grid[i][j]
        }
    }

    p := kmp(s, pattern)
    vis := make([][]int, m)
    for i := range vis {
        vis[i] = make([]int, n)
    }
    for i, j := 0, 0; i < len(p) && j < m * n; j++ {
        if j < p[i] {
            j = p[i]
        } else if j >= p[i] + len(pattern) {
            i++
            if i >= len(p) {
                break
            }
            j = p[i]
        } else if i + 1 < len(p) && j == p[i + 1] {
            i++
        }

        vis[j / n][j % n]++
    }

    for j := range n {
        for i := range m {
            s[j * m + i] = grid[i][j]
        }
    }

    p = kmp(s, pattern)
    for i, j := 0, 0; i < len(p) && j < m * n; j++ {
        if j < p[i] {
            j = p[i]
        } else if j >= p[i] + len(pattern) {
            i++
            if i >= len(p) {
                break
            }
            j = p[i]
        } else if i + 1 < len(p) && j == p[i + 1] {
            i++
        }

        vis[j % m][j / m]++
    }

    for _, v := range vis {
        for _, x := range v {
            if x == 2 {
                ans++
            }
        }
        
    }

    return 
}