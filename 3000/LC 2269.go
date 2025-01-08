func divisorSubstrings(num int, k int) (ans int) {
    m := int(math.Pow10(k))
    for x := num ; x >= m / 10; x /= 10 {
        if x % m > 0 && num % (x % m) == 0 {
            ans++
        }
    }

    return
}

