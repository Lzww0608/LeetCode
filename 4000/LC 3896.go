const N = 100_101
var f [N]bool 
var a []int 

func init() {
    f[1] = true
    for i := 2; i < N; i++ {
        if f[i] {
            continue
        }
        a = append(a, i)
        for j := i * i; j < N; j += i {
            f[j] = true
        }
    }
}

func minOperations(nums []int) (ans int) {
    for i, x := range nums {
        if i & 1 == 0 {
            if f[x] {
                j := sort.SearchInts(a, x)
                ans += a[j] - x
            }
        } else {
            if !f[x] {
                ans++
            }
            if x == 2 {
                ans++
            }
        }
    }

    return
}