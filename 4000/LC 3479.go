func numOfUnplacedFruits(fruits []int, baskets []int) (ans int) {
    n := len(fruits)
    m := int(math.Sqrt(float64(n)))
    section := (n + m - 1) / m 

    mx := make([]int, section)
    for i, x := range baskets {
        mx[i / m] = max(mx[i / m], x)
    }

    for _, x := range fruits {
        unset := 1
        for i, v := range mx {
            if v < x {
                continue
            } 
            mx[i] = 0
            for j := i * m; j < min(n, (i + 1) * m); j++ {
                if baskets[j] >= x && unset == 1 {
                    baskets[j] = 0
                    unset = 0
                }

                mx[i] = max(mx[i], baskets[j])
            }
            
            break
        }
        ans += unset
    }

    return
}