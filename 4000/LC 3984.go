const N int = 1_000_001
const MOD int = 1_000_000_007

var F [N][]int 

func init() {
    for i := 2; i < N; i++ {
        if len(F[i]) != 0 {
            continue
        }
        for j := i; j < N; j += i {
            F[j] = append(F[j], i)
        }
    }
}


func divisibleGame(nums []int) int {
    mx := slices.Max(nums)
    if mx == 1 {
        return MOD - 2
    }

    calc := func(k int) int {
        ans, f := math.MinInt, 0
        for _, x := range nums {
            if x % k != 0 {
                x = -x
            }
            f = max(0, f) + x 
            ans = max(ans, f)
        }

        return ans
    }

    ans, bestk := math.MinInt, 0
    vis := make(map[int]bool)
    for _, v := range nums {
        for _, k := range F[v] {
            if vis[k] {
                continue
            }
            
            vis[k] = true 
            t := calc(k)
            if t > ans {
                ans, bestk = t, k 
            } else if t == ans {
                bestk = min(bestk, k)
            }
                    //fmt.Println(mx, i, j, k)
                
        
        }
    }

    return ans * bestk % MOD
}