func mySqrt(x int) (ans int) {
    l, r := -1, x + 1
    for l < r {
        mid := l + ((r - l) >> 1)
        if mid * mid == x {
            return mid
        } else if mid * mid > x {
            r = mid
        } else {
            l = mid + 1
            ans = mid
        }
    }

    return 
}