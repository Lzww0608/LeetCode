const N int = 1_001

var F [N]int 

func init() {
    for i := 2; i < N; i++ {
        f := true
        for j := 2; j * j <= i; j++ {
            if i % j == 0 {
                f = false 
                break
            }
        }

        if f {
            F[i] = F[i - 1] + i
        } else {
            F[i] = F[i - 1]
        }
    }
}

func sumOfPrimesInRange(x int) int {
    y := 0
    for t := x; t > 0; t /= 10 {
        y = y * 10 + t % 10
    }

    x, y = min(x, y), max(x, y)
    return F[y] - F[x - 1]
}