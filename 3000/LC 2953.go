
import "math/rand"
const N int = 26
func countCompleteSubstrings(s string, count int) (ans int) {
    n := len(s)
    base := make([]uint64, 26)
    cnt := make([]int, N)
    cntH := map[uint64]int{0: 1}
    h := make([]uint64, n + 1)
    rem := make([]int, N)


    for i := 0; i < N; i++ {
        base[i] = rand.Uint64()
    }

    l := 0
    for i := range s {
        if i > 0 && abs(int(s[i]-'a')-int(s[i-1]-'a')) > 2 {
            cntH = map[uint64]int{0: 1}
            cnt = make([]int, N)
            h[i] = 0
            rem = make([]int, N)
            l = i
        }
        x := int(s[i] - 'a')
        cnt[x]++
        for cnt[x] > count {
            cnt[int(s[l]-'a')]--
            cntH[h[l]]--
            l++
        }
        h[i+1] = h[i] + uint64((rem[x] + 1) % count - (rem[x]) % count) * base[x]
        ans += cntH[h[i+1]]
        rem[x]++
        cntH[h[i+1]]++
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}