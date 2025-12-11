func longestBalanced(nums []int) int {
    n := len(nums)
    for d := n; d > 0; d-- {
        cnt := make(map[int]int)
        sum := 0
        for i, x := range nums {
            if cnt[x]++; cnt[x] == 1 {
                if x & 1 == 0 {
                    sum++
                } else {
                    sum--
                }
            }

            if i >= d - 1 {
                if sum == 0 {
                    return d
                }
                y := nums[i - d + 1]
                if cnt[y]--; cnt[y] == 0 {
                    if y & 1 == 0 {
                        sum--
                    } else {
                        sum++
                    }
                }
            }
        }
    }

    return 0
}