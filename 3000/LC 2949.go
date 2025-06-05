const AEIOU_MASK int = (1 << 0) + (1 << 4) + (1 << 8) + (1 << 14) + (1 << 20)

func sqrt(n int) int {
    res := 1
    for i := 2; i * i <= n; i++ {
        j := i * i 
        for n % j == 0 {
            res *= i 
            n /= j
        }
        if n % i == 0 {
            res *= i 
            n /= i
        }
    }

    if n > 1 {
        res *= n
    }

    return res
}

func beautifulSubstrings(s string, k int) (ans int64) {
    k = sqrt(k * 4)

    type pair struct {
        i, sum int
    }

    cnt := make(map[pair]int)
    cnt[pair{k - 1, 0}] = 1 
    sum := 0
    for i, c := range s {
        bit := (AEIOU_MASK >> (c - 'a')) & 1
        // 1: 1, 0: -1
        sum += bit * 2 - 1 
        p := pair{i % k, sum}
        ans += int64(cnt[p])
        cnt[p]++
    }

    return
}