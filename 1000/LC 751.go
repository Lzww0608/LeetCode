func ipToCIDR(ip string, n int) (ans []string) {
    s := strings.Split(ip, ".")
    num := 0
    for _, t := range s {
        x, _ := strconv.Atoi(t)
        num = num * 256 + x
    }

    for n > 0 {
        t := num & -num
        k := max(33 - bits.Len(uint(t)), 33 - bits.Len(uint(n)))
        if t == 0 {
            k = 33 - bits.Len(uint(n))
        }
        ans = append(ans, ipToInt(num) + "/" + strconv.Itoa(k))
        num += 1 << (32 - k)
        n -= 1 << (32 - k)
    }

    return
}

func ipToInt(x int) string {
    s := make([]string, 4)
    for i := 3; i >= 0; i-- {
        t := strconv.Itoa(x % 256)
        x >>= 8
        s[i] = t
    }
    return strings.Join(s, ".")
}
