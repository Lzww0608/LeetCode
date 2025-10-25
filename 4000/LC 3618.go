const N int = 100_001

var isPrime [N]int 

func init() {
    isPrime[0], isPrime[1] = 1, 1
    for i := 2; i * i < N; i++ {
        if isPrime[i] == 1 {
            continue
        }
        for j := i * i; j < N; j += i {
            isPrime[j] = 1
        }
    }
}

func splitArray(nums []int) int64 {
    ans := 0
    for i, x := range nums {
        if isPrime[i] == 0 {
            ans += x
        } else {
            ans -= x
        }
    }
    
    if ans < 0 {
        return -int64(ans)
    }
    return int64(ans)
}