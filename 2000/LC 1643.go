func kthSmallestPath(destination []int, k int) string {
    h, v := destination[1], destination[0]

    comb := make([][]int, h + v)
    for i := range comb {
        comb[i] = make([]int, h)
    }
    comb[0][0] = 1

    for i := 1; i < h + v; i++ {
        comb[i][0] = 1
        for j := 1; j <= i && j < h; j++ {
            comb[i][j] = comb[i-1][j-1] + comb[i-1][j];
        }
    }

    ans := []byte{}
    for i, mx := 0, h + v; i < mx; i++ {
        if h > 0 {
            x := comb[h+v-1][h-1];
            if k > x {
                ans = append(ans, 'V')
                v--
                k -= x
            } else {
                ans = append(ans, 'H')
                h--
            }
        } else {
            ans = append(ans, 'V')
            v--
        }
    }

    return string(ans)
}