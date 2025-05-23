func minimizeError(prices []string, target int) string {
    s := []float64{}
    for _, price := range prices {
        t := strings.Split(price, ".")
        x, _ := strconv.Atoi(t[0])
        target -= x
        y, _ := strconv.Atoi(t[1])
        if y != 0 {
            s = append(s, float64(y) / 1000.0)
        }
        
    }

    n := len(s)
    if target < 0 || target > n {
        return "-1"
    }

    sort.Slice(s, func(i, j int) bool {
        return s[i] > s[j]
    })

    var ans float64
    for i := 0; i < n; i++ {
        if target > 0 {
            target--
            ans += 1.0 - s[i]
        } else {
            ans += s[i]
        }
    } 

    return fmt.Sprintf("%.3f", ans)
}