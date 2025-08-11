func countEven(num int) int {
    if num & 1 == 0 {
        return num / 2 - calc(num)
    }
    return num / 2
}

func calc(x int) int {
    cnt := 0
    for x > 0 {
        cnt += x % 10
        x /= 10
    }
    
    return cnt & 1
}