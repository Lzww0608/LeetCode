func findGameWinner(n int) bool {
    if n == 1 {
        return false
    }

    sg := make([]int, n)
    sg[0], sg[1] = 0, 1
    for i := 2; i < n; i++ {
        sg[i] = 1 + (sg[i-1] ^ sg[i-2])
    } 
    
    return (sg[n-1] ^ sg[n-2]) > 0
}