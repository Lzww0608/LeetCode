const N = 33
func countPrimeSetBits(left int, right int) (ans int) {
    f := make([]int, N)
    f[0], f[1] = 1, 1
    for i := 2; i < N; i++ {
        if f[i] == 1 {
            continue
        }
        for j := i * i; j < N; j += i {
            f[j] = 1
        }
    }
    

    for i := left; i <= right; i++ {
        x := bits.OnesCount32(uint32(i))
        if f[x] == 0 {
            ans++
        }
    }

    return 
}