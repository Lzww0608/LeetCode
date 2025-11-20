func sumOfEncryptedInt(nums []int) (ans int) {
    for _, x := range nums {
        ans += calc(x)
    }

    return
}

func calc(x int) int {
    mx := 0
    for t := x; t > 0; t /= 10 {
        mx = max(mx, t % 10)
    }

    res := 0
    for x > 0 {
        res = res * 10 + mx 
        x /= 10
    }

    return res
}