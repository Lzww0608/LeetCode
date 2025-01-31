
import "math/bits"
func getMaxFunctionValue(receiver []int, k int64) int64 {
    n, m := len(receiver), bits.Len(uint(k))
    fa := make([][]pair, n)
    for i, x := range receiver {
        fa[i] = make([]pair, m)
        fa[i][0] = pair{x, x}
    }

    for i := 0; i + 1 < m; i++ {
        for j := range fa {
            f := fa[j][i]
            ff := fa[f.fa][i]
            fa[j][i+1] = pair{ff.fa, f.sum + ff.sum}
        }
    }

    ans := 0
    for i := 0; i < n; i++ {
        j := i 
        sum := i 
        for t := uint(k); t > 0; t &= t - 1 {
            f := fa[j][bits.TrailingZeros(t)]
            sum += f.sum
            j = f.fa
        }
        ans = max(ans, sum)
    }

    return int64(ans)
}

type pair struct {
    fa, sum int
}