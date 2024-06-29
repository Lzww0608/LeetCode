func lemonadeChange(bills []int) bool {
    five, ten := 0, 0
    for _, x := range bills {
        if x == 5 {
            five++
        } else if x == 10 {
            if five > 0 {
                five--
                ten++
            } else {
                return false;
            }
        } else {
            if ten > 0 && five > 0 {
                ten--
                five--
            } else if five > 2 {
                five -= 3
            } else {
                return false
            }
        }
    }
    return true
}
