const N int = 51

var comb [N][N]int64 

func init() {
    for i := range N {
        comb[i][0] = 1
        for j := 1; j <= i; j++ {
            comb[i][j] = comb[i - 1][j - 1] + comb[i - 1][j]
        }
    }
}

func nthSmallest(n int64, k int) int64 {
    ans := 0
    for i := 49; i >= 0; i-- {
        x := comb[i][k]
        if n > x {
            n -= x 
            k--
            ans |= 1 << i
        }
    }

    return int64(ans)
}