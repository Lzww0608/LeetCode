const MOD int = 1_000_000_007
const N int = 100_005
var factor [N]int

func init() {
    for i := 2; i < N; i++ {
        if factor[i] != 0 {
            continue
        }
        for j := i; j < N; j += i {
            factor[j] += 1
        }
    }
}

func maximumScore(nums []int, k int) (ans int) {
    ans = 1
    n := len(nums)
    left := make([]int, n)
    right := make([]int, n)
    for i := 0; i < n; i++ {
        left[i] = -1
        right[i] = n
    }

    st1 := []int{}
    st2 := []int{}
    for i := 0; i < n; i++ {
        x, y := factor[nums[i]], factor[nums[n-1-i]]
        for len(st1) > 0 && x > factor[nums[st1[len(st1)-1]]] {
            st1 = st1[:len(st1)-1]
        }
        if len(st1) > 0 {
            left[i] = st1[len(st1)-1]
        }
        st1 = append(st1, i)

        for len(st2) > 0 && y >= factor[nums[st2[len(st2)-1]]] {
            st2 = st2[:len(st2)-1]
        }
        if len(st2) > 0 {
            right[n-1-i] = st2[len(st2)-1]
        }
        st2 = append(st2, n - 1 - i)
    }

    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return nums[id[i]] > nums[id[j]]
    })

    for _, i := range id {
        sum := (i - left[i]) * (right[i] - i) 
        if sum >= k {
            ans = (ans * quickPow(nums[i], k)) % MOD
            break
        }
        ans = (ans * quickPow(nums[i], sum)) % MOD
        k -= sum
    }

    return
}


func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = (res * a) % MOD
        }
        a = (a * a) % MOD
        r >>= 1
    }

    return res
}