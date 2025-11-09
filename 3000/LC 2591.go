func distMoney(money int, children int) int {
    if money < children {
        return -1
    }

    if money > children * 8 {
        return children - 1
    }

    //x * 8 + (children - x) <= money
    // x * 7 <= money - children
    x := (money - children) / 7 
    if money == (children - 1) * 8 + 4 {
        return x - 1
    }

    return x
}