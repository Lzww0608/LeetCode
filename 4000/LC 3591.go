func checkPrimeFrequency(nums []int) bool {
    cnt := make(map[int]int)
    for _, x := range nums {
        cnt[x]++
    }
    for _, v := range cnt {
        if check(v) {
            return true
        }
    }

    return false
}

func check(x int) bool {
    if x == 1 {
        return false
    }
    for i := 2; i * i <= x; i++ {
        if x % i == 0 {
            return false 
        }
    }

    return true
}