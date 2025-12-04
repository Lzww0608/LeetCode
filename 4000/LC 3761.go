func minMirrorPairDistance(nums []int) int {
    p := make(map[int]int)
    ans := len(nums) + 1
    for i, x := range nums {
        y := reverse(x)
        if j, ok := p[x]; ok {
            ans = min(ans, i - j)
        }
        p[y] = i
    }
    
    if ans > len(nums) {
        return -1
    }

    return ans
}

func reverse(x int) int {
    res := 0
    for x > 0 {
        res = res * 10 + x % 10
        x /= 10
    }

    return res
}