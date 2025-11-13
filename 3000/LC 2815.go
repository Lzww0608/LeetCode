
func maxSum(nums []int) int {
    ans := -1
    for i := range nums {
        for j := i + 1; j < len(nums); j++ {
            x := findMax(nums[i])
            y := findMax(nums[j])
            if x == y {
                ans = max(ans, nums[i] + nums[j])
            }
        }
    }
    const N int = 10
func maxSum(nums []int) int {
    ans := -1
    a := [N]int{}
    for _, x := range nums {
        y := findMax(x)
        if a[y] == 0 {
            a[y] = x
        } else {
            ans = max(ans, x + a[y])
            a[y] = max(a[y], x)
        }
    }
    
    return ans
}

func findMax(x int) (ans int) {
    for x > 0 {
        ans = max(ans, x % 10)
        x /= 10
    }

    return
}
    return ans
}

func findMax(x int) (ans int) {
    for x > 0 {
        ans = max(ans, x % 10)
        x /= 10
    }

    return
}