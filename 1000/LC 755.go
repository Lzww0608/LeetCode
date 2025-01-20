func pourWater(heights []int, volume int, k int) []int {
    n := len(heights)
next:
    for ; volume > 0; volume-- {
        for d := -1; d < 2; d += 2 {
            i := k
            pos := i 
            for i >= 0 && i < n {
                if heights[i] < heights[pos] {
                    pos = i 
                } else if heights[i] > heights[pos] {
                    break
                }
                i += d
            }
            if pos != k {
                heights[pos]++
                continue next
            }
        }
        heights[k]++
    }

    return heights
}