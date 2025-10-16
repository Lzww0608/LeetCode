func minSwaps(nums []int) int {
    n := len(nums)
    odd := []int{}
    even := []int{}
    for i, x := range nums {
        if x & 1 == 1 {
            odd = append(odd, i)
        } else {
            even = append(even, i)
        }
    }

    if len(odd) - len(even) > 1 || len(even) - len(odd) > 1 {
        return -1
    }

    sum1, sum2 := 0, 0 
    for i, j := 0, 0; i < len(odd) || j < n; j += 2 {
        if i >= len(odd) {
            sum1 = math.MaxInt32
            break
        }
        sum1 += abs(odd[i] - j)
        i++
    }

    for i, j := 0, 1; i < len(odd) || j < n; j += 2 {
        if i >= len(odd) {
            sum2 = math.MaxInt32
            break
        }func minSwaps(nums []int) int {
    n := len(nums)
    odd := []int{}
    even := []int{}
    for i, x := range nums {
        if x & 1 == 1 {
            odd = append(odd, i)
        } else {
            even = append(even, i)
        }
    }

    if len(odd) - len(even) > 1 || len(even) - len(odd) > 1 {
        return -1
    }

    sum1, sum2 := 0, 0 
    for i, j := 0, 0; j < n; j += 2 {
        if i >= len(odd) || len(odd) != (n + 1) / 2 {
            sum1 = math.MaxInt32
            break
        }
        sum1 += abs(odd[i] - j)
        i++
    }

    for i, j := 0, 1; j < n; j += 2 {
        if i >= len(odd) || len(odd) != n / 2 {
            sum2 = math.MaxInt32
            break
        }
        sum2 += abs(odd[i] - j)
        i++
    }

    return min(sum1, sum2)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}
        sum2 += abs(odd[i] - j)
        i++
    }

    return min(sum1, sum2)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}