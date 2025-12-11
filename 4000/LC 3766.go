const N int = 5001

var p [N]int 

func check(x int) bool {
    s := strconv.FormatInt(int64(x), 2)
    n := len(s)
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
        if s[i] != s[j] {
            return false
        }
    }

    return true
}

func init() {
    pre := 0
    for i := 1; i < N; i++ {
        p[i] = N
        if check(i) {
            pre = 0
        } else {
            pre++
        }

        p[i] = pre
    }

    pre = N 
    for i := N - 1; i > 0; i-- {
        if p[i] == 0 {
            pre = 0
        } else {
            pre++
        }

        p[i] = min(p[i], pre)
    }
}

func minOperations(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)

    for i, x := range nums {
        ans[i] = p[x]
    }

    return ans
}