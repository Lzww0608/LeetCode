func kthDigit(k int64) int {
    if k <= 9 {
        return int(k)
    }

    k -= 9
    var d int64 = 2
    var count int64 = 90

    for {
        digits := d * count
        if k > digits {
            k -= digits
            d++
            count *= 10
            continue
        }

        offset := k - 1
        numIndex := offset / d
        digitIndex := offset % d

        start := count / 9
        num := start + numIndex

        block := num / 10
        pos := num % 10

        if block%2 == 1 {
            num = block*10 + 9 - pos
        }

        s := strconv.FormatInt(num, 10)
        return int(s[digitIndex] - '0')
    }
}