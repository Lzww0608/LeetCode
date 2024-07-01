func maximumPrimeDifference(nums []int) int {
    f := make([]bool, 101)
    f[1] = true
    for i := 2; i * i < len(f); i++ {
        if f[i] {
            continue
        }
         for j := i * i; j < len(f); j += i {
            f[j] = true
        }
    }


    i, j := 0, len(nums) - 1
    for f[nums[i]] {
        i++
    }

    for f[nums[j]] {
        j--
    }

    return j - i
}

