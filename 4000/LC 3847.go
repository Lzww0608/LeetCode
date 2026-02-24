func scoreDifference(nums []int) (ans int) {
    f := true
    for i, x := range nums {
        if x & 1 == 1 {
            f = !f
        }
        if (i + 1) % 6 == 0 {
            f = !f
        }

        if f {
            ans += x
        } else {
            ans -= x
        }
    }

    return
}