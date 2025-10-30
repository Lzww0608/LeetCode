func maxBalancedShipments(weight []int) (ans int) {
    mn := math.MaxInt32
    for _, x := range weight {
        if mn != math.MaxInt32 && mn > x {
            ans++ 
            mn = math.MaxInt32
        } else {
            if mn == math.MaxInt32 {
                mn = x
            } else if x > mn {
                mn = x
            }
        }
    }

    return
}