const N int = 100_005

var isPrimes [N]int

var s [N][]int

func init() {
    for i := 2; i < N; i++ {
        if isPrimes[i] == 0 {
            for j := i * 2; j < N; j += i {
                s[j] = append(s[j], i)
                isPrimes[j] = 1
            }
        }
    }
}
    

func canTraverseAllPairs(nums []int) bool {
    m := make(map[int]bool)
    for _, x := range nums {
        m[x] = true
    }
    fa := make([]int, N)
    for i := range fa {
        fa[i] = i 
    }

    var find func(int) int 
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }

        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if ry > rx {
            rx, ry = ry, rx
        }
        if rx != ry {
            fa[rx] = ry
        }
    }

    for _, x := range nums {
        for _, y := range s[x] {
            merge(x, y)
        }
    }

    root := find(nums[0])
    //fmt.Println(root)
    for _, x := range nums[1:] {
        if find(x) != root || root == 1 {
            return false
        }
    }

    return true
}