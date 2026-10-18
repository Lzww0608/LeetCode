const N int = 1_001

var a []int 
var isPrime [N]bool

func init() {
    for i := 2; i < N; i++ {
        if isPrime[i] {
            continue
        }
        a = append(a, i)
        for j := i * i; j < N; j += i {
            isPrime[j] = true
        }
    }
}

func maxPrimeSubset(n int, s int) (ans []int) {
    m := sort.SearchInts(a, n + 1)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, s + 1)
        for j := range f[i] {
            f[i][j] = math.MaxInt32
        }
    }
    f[m][0] = 0

    for i := m - 1; i >= 0; i-- {
        x := a[i]
        for j := range s + 1 {
            if j < x {
                f[i][j] = f[i + 1][j]
            } else {
                f[i][j] = min(f[i + 1][j], f[i + 1][j - x] + 1)
            }
        }
        
    }

    j := s 
    for j > 0 && f[0][j] == math.MaxInt32 {
        j--
    }

    for i, x := range a[:m] {
        if j >= x && f[i + 1][j - x] + 1 <= f[i][j] {
            j -= x
            ans = append(ans, x)
        }
    }

    return
}