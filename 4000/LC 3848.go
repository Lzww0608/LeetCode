func isDigitorialPermutation(n int) bool {
    m := map[int]int {
        0: 1,
        1: 1,
        2: 2,
        3: 6,
        4: 24,
        5: 120,
        6: 720,
        7: 5040,
        8: 40320,
        9: 362880,
    }
    sum := 0
    cnt := [10]int{}
    for x := n; x > 0; x /= 10 {
        sum += m[x % 10]
        cnt[x % 10]++
    }

    for sum > 0 {
        cnt[sum % 10]--
        sum /= 10
    }

    for _, x := range cnt {
        if x != 0 {
            return false
        }
    }

    return true
}