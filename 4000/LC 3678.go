func smallestAbsent(nums []int) (ans int) {
    sum := 0
    cnt := make(map[int]bool)
    for _, x := range nums {
        sum += x
        cnt[x] = true
    }

    avg := max(0, sum / len(nums))
    for ans = avg + 1; cnt[ans]; ans++ {

    }

    return
}