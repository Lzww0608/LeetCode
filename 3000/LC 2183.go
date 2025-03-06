const N int = 100_005
var factor [N][]int

func init() {
    for i := 1; i < N; i++ {
        for j := i; j < N; j += i {
            factor[j] = append(factor[j], i)
        }
    }
}

func countPairs(nums []int, k int) int64 {
    ans := 0
    cnt := make(map[int]int)
    for _, x := range nums {
        t := k / gcd(k, x)
        ans += cnt[t]
        for _, y := range factor[x] {
            cnt[y]++
        }
    }

    return int64(ans)
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y 
    }
    return x 
}