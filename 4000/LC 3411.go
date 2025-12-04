func maxLength(nums []int) int {
    ans := 2

    prod := 1
    for l, r := 0, 0; r < len(nums); r++ {
        for gcd(prod, nums[r]) > 1 {
            prod /= nums[l]
            l++
        }

        prod *= nums[r]
        ans = max(ans, r - l + 1)
    }

    return ans
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b 
    }

    return a
}