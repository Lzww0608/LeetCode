func findShortestSubArray(nums []int) int {
    type tuple struct {
        k, i, j int
    }

    m := map[int]tuple{}
    for i, x := range nums {
        if v, ok := m[x]; !ok {
            m[x] = tuple{1, i, i}
        } else {
            v.k++
            v.j = i
            m[x] = v
        }
    }
    ans, cnt := math.MaxInt32, 0
    for _, v := range m {
        if v.k > cnt {
            cnt = v.k
            ans = v.j - v.i + 1
        } else if v.k == cnt {
            ans = min(ans, v.j - v.i + 1)
        }
    }
    if ans == math.MaxInt32 {
        return 0
    }

    return ans
}