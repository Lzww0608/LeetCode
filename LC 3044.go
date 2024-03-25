func mostFrequentPrime(mat [][]int) int {
    mp := map[int]int{}
    n, m := len(mat), len(mat[0])
    f := make([]int, int(math.Pow(10, float64(max(m, n)))))

    for i := 2; i < len(f); i++ {
        if f[i] == 1 {
            continue
        } 
        for j := int64(i) * int64(i); j < int64(len(f)); j += int64(i) {
            f[int(j)] = 1
        }
    }
   
    dir := [8][2]int{
        {0, 1},
        {0, -1},
        {1, 0},
        {1, 1},
        {1, -1},
        {-1, 0},
        {-1, 1},
        {-1, -1},
    }

    cal := func(i, j int) {
        for k := range dir {
            t := 0
            x, y := dir[k][0], dir[k][1]
            a, b := i, j
            for a >= 0 && a < n && b >= 0 && b < m {
                t = t * 10 + mat[a][b]
                if t > 10 && f[t] == 0 {
                    mp[t]++
                }
                a, b = a + x, b + y
            }
            
        }
    }

    for i := range mat {
        for j := range mat[i] {
            cal(i, j)
        }
    }

    ans, cnt := -1, 0
    for k, v := range mp {
        //fmt.Println(k, v)
        if v > cnt || (v == cnt && k > ans) {
            ans, cnt = k, v
        }
    } 
    return ans
}