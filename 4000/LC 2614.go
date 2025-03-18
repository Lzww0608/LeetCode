func diagonalPrime(nums [][]int) (ans int) {
    n := len(nums)
    for i := 0; i < n; i++ {
        if ans < nums[i][i] && check(nums[i][i]) {
            ans = nums[i][i]
        }

        if ans < nums[i][n-i-1] && check(nums[i][n-i-1]) {
            ans = nums[i][n-i-1]
        }
    }


    return 
}

func check(x int) bool {
    if x == 1 {
        return false
    }
    for i := 2; i * i <= x; i++ {
        if x % i == 0 {
            return false
        }
    }

    return true
}

