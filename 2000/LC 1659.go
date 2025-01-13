func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
    p := int(math.Pow(3, float64(n-1)))
    h := [3][3]int{{0, 0, 0}, {0, -60, -10}, {0, -10, 40}}
    memo := make([][][][]int, m * n)
    for i := range memo {
		memo[i] = make([][][]int, p*3)
		for j := range memo[i] {
			memo[i][j] = make([][]int, introvertsCount+1)
			for k := range memo[i][j] {
				memo[i][j][k] = make([]int, extrovertsCount+1)
				for l := range memo[i][j][k] {
					memo[i][j][k][l] = -1
				}
			}
		}
	}

    var dfs func(int, int, int, int) int
    dfs = func(pos, pre, ic, ec int) (res int) {
        if pos == m * n || (ic == 0 && ec == 0) {
            return 0
        }
        ptr := &memo[pos][pre][ic][ec]
        if *ptr != -1 {
            return *ptr 
        }

        up := pre / p
        left := pre % 3
        if pos % n == 0 {
            left = 0
        }
        for i := 0; i < 3; i++ {
            if i == 1 && ic == 0 || i == 2 && ec == 0 {
                continue
            }
            cur := pre % p * 3 + i
            cur_ic, cur_ec := ic, ec 
            c := 0
            if i == 1 {
                cur_ic--
                c = 120
            } else if i == 2 {
                cur_ec--
                c = 40
            }
            a := h[up][i] + h[left][i]
            b := dfs(pos + 1, cur, cur_ic, cur_ec)
            res = max(res, a + b +c)
        }
        *ptr = res 
        return
    }

    return dfs(0, 0, introvertsCount, extrovertsCount)
}