func averageValue(nums []int) int {
    cnt, sum := 0, 0
    for _, x := range nums {
        if x % 3 == 0 && x & 1 == 0 {
            sum += x 
            cnt++
        }
    }

    if cnt == 0 {
        return 0
    }

    return sum / cnt 
}