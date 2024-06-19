func minimumSize(nums []int, maxOperations int) (ans int) {
    mx := nums[0]
    for _, x := range nums {
        mx = max(mx, x)
    }

    l, r := 1, mx

    for l <= r {
        mid := l + ((r - l) >> 1)
        cost := 0
        
        for _, x := range nums {
            cost += (x - 1) / mid 
        }

        if cost <= maxOperations {
            ans = mid
            r = mid - 1
        } else {
            l = mid + 1
        }
    }

    return 
}
