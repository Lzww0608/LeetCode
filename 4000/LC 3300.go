func minElement(nums []int) (ans int) {
    ans = math.MaxInt32
    for _, x := range nums {
        ans = min(ans, calc(x))
    }

    return
}

func calc(x int) (res int) {
    for x > 0 {
        res += x % 10
        x /= 10
    }

    return
}