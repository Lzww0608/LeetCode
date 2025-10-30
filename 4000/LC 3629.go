const N int = 1_000_001

var primes [N][]int 

func init() {
    for i := 2; i < N; i++ {
        if len(primes[i]) != 0 {
            continue
        }

        for j := i; j < N; j += i {
            primes[j] = append(primes[j], i)
        }
    }
}

func minJumps(nums []int) int {
    n := len(nums)
    m := make(map[int][]int)
    
    for i, x := range nums {
        for _, y := range primes[x] {
            m[y] = append(m[y], i) 
        }
    }

    vis := make([]bool, n)
    ans := 0
    vis[0] = true 
    q := []int{0}

    for {
        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            q = q[1:]
            if cur == n - 1 {
                return ans
            }
            if !vis[cur + 1] {
                vis[cur + 1] = true 
                q = append(q, cur + 1)
            }
            if cur - 1 >= 0 && !vis[cur - 1] {
                vis[cur - 1] = true 
                q = append(q, cur - 1)
            }

            for _, i := range m[nums[cur]] {
                if !vis[i] {
                    vis[i] = true 
                    q = append(q, i)
                }
            }
            m[nums[cur]] = []int{}
        }
        ans++
    }

    return ans
}