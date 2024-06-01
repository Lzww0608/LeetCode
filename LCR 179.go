func twoSum(price []int, target int) []int {
    i, j := 0, len(price) - 1
    for i < j {
        s := price[i] + price[j]
        if s < target {
            i++
        } else if s > target {
            j--
        } else {
            return []int{price[i], price[j]}
        }
    }

    return nil
}