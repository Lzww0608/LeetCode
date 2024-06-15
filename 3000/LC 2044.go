func countMaxOrSubsets(nums []int) (ans int) {
    maxOr := 0
    n := len(nums)
    f := make([]int, 1 << n)
    for i, x := range nums {
        maxOr |= x
        f[1<<i] = x
    }
    

    for i := 1; i < (1 << n); i++ {
        x := i & (-i)
        if x ^ i != 0 {
            f[i] = f[i ^ x] | f[x]
        }
        if f[i] == maxOr {
            ans++
        }
    }



    return
}