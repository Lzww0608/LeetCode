func distinctPrimeFactors(nums []int) int {
    m := make(map[int]struct{})

    for _, x := range nums {
        for i := 2; i * i <= x; i++ {
            if x % i == 0 {
                m[i] = struct{}{}
                for x /= i; x % i == 0; x /= i {}
            }
        }
        if x > 1 {
            m[x] = struct{}{}
        }
    }

    return len(m)
}