func smallestNumber(n int, t int) int {
    x := n 
    for calc(x) % t != 0 {
        x++
    }

    return x
}

func calc(x int) (res int) {
    res = 1
    for x > 0 {
        res *= x % 10 
        x /= 10
    }

    return
}