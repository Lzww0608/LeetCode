func longestBalanced(s string) int {
    return max(checkSingle(s), checkDouble(s), checkTriple(s))
}

func checkSingle(s string) (ans int) {
    for _, c := range "abc" {
        cnt := 0
        for _, t := range s {
            if t != c {
                cnt = 0
            } else {
                cnt++
                ans = max(ans, cnt)
            }
        }
    }

    return 
}

func checkDouble(s string) (ans int) {
	n := len(s)

	pos := make([]int, 2*n+1)
	for i := range pos {
		pos[i] = -2
	}

	history := make([]int, 0, n)

	pairs := []struct{ a, b byte }{
		{'a', 'b'}, {'b', 'c'}, {'a', 'c'},
	}

	for _, p := range pairs {
		a, b := p.a, p.b
		
		for _, idx := range history {
			pos[idx] = -2
		}
		history = history[:0]

		diff := 0
		pos[n] = -1
		history = append(history, n)

		for i := 0; i < n; i++ {
			if s[i] == a {
				diff++
			} else if s[i] == b {
				diff--
			} else {
				for _, idx := range history {
					pos[idx] = -2
				}
				history = history[:0]
				
				diff = 0
				pos[n] = i
				history = append(history, n)
				continue
			}

			idx := diff + n
			if pos[idx] != -2 {
				ans = max(ans, i-pos[idx])
			} else {
				pos[idx] = i
				history = append(history, idx)
			}
		}
	}
	return
}

func checkTriple(s string) (ans int) {
    type pair struct {
        x, y int
    }

    m := make(map[pair]int)
    m[pair{0, 0}] = -1
    a, b, c := 0, 0, 0
    for i := range s {
        if s[i] == 'a' {
            a++
        } else if s[i] == 'b' {
            b++
        } else {
            c++
        }

        p := pair{a - b, a - c}
        if j, ok := m[p]; ok {
            ans = max(ans, i - j)
        } else {
            m[p] = i
        }
    }

    return
}