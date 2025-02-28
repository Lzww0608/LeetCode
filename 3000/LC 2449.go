func makeSimilar(nums []int, target []int) int64 {
    n := len(nums)
    odd1 := make([]int, 0, n)
    odd2 := make([]int, 0, n)
    even1 := make([]int, 0, n)
    even2 := make([]int, 0, n)
    for _, x := range nums {
        if x & 1 == 1 {
            odd1 = append(odd1, x)
        } else {
            even1 = append(even1, x)
        }
    }

    for _, x := range target {
        if x & 1 == 1 {
            odd2 = append(odd2, x)
        } else {
            even2 = append(even2, x)
        }
    }

    sort.Ints(odd1)
    sort.Ints(odd2)
    sort.Ints(even1)
    sort.Ints(even2)

    ans := 0
    for i := 0; i < len(odd1); i++ {
        ans += abs(odd2[i] - odd1[i])
    }
    for i := 0; i < len(even1); i++ {
        ans += abs(even2[i] - even1[i]) 
    }

    return int64(ans) / 4
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}