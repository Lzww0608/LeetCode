func countValidSelections(nums []int) (ans int) {
    sum := 0
    for _, x := range nums {
        sum += x 
    }

    pre := 0
    for _, x := range nums {
        if x != 0 {
            pre += x
        } else {
            if pre == sum - pre {
                ans += 2
            } else if pre * 2 + 1 == sum || pre * 2 - 1 == sum {
                ans++
            }
        }
    }

    return
}