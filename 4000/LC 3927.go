const N int = 100_001

func minArraySum(nums []int) int64 {
    vis := make([]bool, N)
    mx := 0
    for _, x := range nums {
        vis[x] = true
        mx = max(mx, x)
    }    

    n := len(nums)
    if vis[1] {
        return int64(n)
    }

    p := make([]int, mx + 1)
    for i := 2; i <= mx; i++ {
        if p[i] != 0 || !vis[i] {
            continue
        }

        for j := i; j <= mx; j += i {
            if p[j] == 0 {
                p[j] = i
            }
            
        }
    }

    ans := 0
    for _, x := range nums {
        ans += p[x]
    }

    return int64(ans)
}