func maxGcdSum(nums []int, k int) int64 {
    n := len(nums)
    ans := 0
    type pair struct {
        i, x int 
    }
    pre := make([]int, n + 1)

    s := []pair{}
    for i, x := range nums {
        pre[i+1] = pre[i] + x
        s = append(s, pair{i, x})

        j := 0
        for k, v := range s {
            s[k].x = gcd(v.x, x)
            if s[k].x != s[j].x {
                j++
                s[j] = s[k]
            } else {
                continue
            }
        }

        s = s[:j+1]
        for _, v := range s {
            if i - v.i + 1 >= k {
                ans = max(ans, v.x * (pre[i+1] - pre[v.i]))
            }
        }
    }

    return int64(ans)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}