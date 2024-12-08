func findValidSplit(nums []int) int {
    n := len(nums)
    left := make(map[int]int)
    right := make([]int, n)

    f := func(x, i int) {
        if v, ok := left[x]; ok {
            right[v] = i 
        } else {
            left[x] = i
        }
    }

    for i, x := range nums {
        for j := 2; j * j <= x; j++ {
            if x % j == 0 {
                f(j, i)
                for x /= j; x % j == 0; x /= j {}
            }
        }

        if x > 1 {
            f(x, i)
        }
    }

    mx := 0
    for i, r := range right {
        if i > mx {
            return mx
        } else {
            mx = max(mx, r)
        }
    }

    return -1
}