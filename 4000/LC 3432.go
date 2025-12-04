func countPartitions(nums []int) (ans int) {
    sum := 0
    for _, x := range nums {
        sum += x
    }

    cur := 0
    for _, x := range nums[:len(nums) - 1] {
        cur += x 
        if (sum - cur * 2) & 1 == 0 {
            ans++
        }
    }

    return
}
