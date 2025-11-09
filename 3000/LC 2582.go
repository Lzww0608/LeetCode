func passThePillow(n int, time int) int {
    if (time / (n - 1)) & 1 == 1 {
        time %= n - 1
        return n - time
    } 

    time %= n - 1
    return time + 1
}