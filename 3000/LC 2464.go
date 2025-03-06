func validSubarraySplit(nums []int) int {
    n := len(nums)
    if nums[0] == 1 || nums[n-1] == 1 {
        return -1
    }
    f := make([]int, n + 1)
    for i := range f {
        f[i] = math.MaxInt32
    }
    f[0] = 0

    for i, x := range nums {
        if x == 1 {
            continue
        }
        for j := i; j < n; j++ {
            if gcd(x, nums[j]) {
                f[j+1] = min(f[j+1], f[i] + 1)
            }
        }
    }

    if f[n] == math.MaxInt32 {
        return -1
    }

    return f[n]
}

func gcd(x, y int) bool {
    for y != 0 {
        x, y = y, x % y 
    }
    return x > 1
}