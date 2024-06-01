func maxSatisfied(customers []int, grumpy []int, minutes int) (ans int) {

    mx, cur := 0, 0
    for i, x := range grumpy {
        if x == 0 {
            ans += customers[i]
            customers[i] = 0
        } else {
            cur += customers[i]
        }
        
        if i >= minutes {
            cur -= customers[i - minutes]
        }
        mx = max(mx, cur)
    }

    ans += mx

    return
}
