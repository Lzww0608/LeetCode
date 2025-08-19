const N int = 100_001

var F []int = make([]int, 0, N)

func init() {
	for base := 1; base <= 10000; base *= 10 {
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			F = append(F, x)
		}
		if base <= 1000 {
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				F = append(F, x)
			}
		}
	}
	F = append(F, 1_000_000_001) // 哨兵，防止下标越界
}

func minimumCost(nums []int) int64 {
	sort.Ints(nums)
    sort.Ints(F)

    calc := func(x int) (res int) {
        for _, y := range nums {
            res += abs(y - x)
        }

        return 
    }

    n := len(nums)
    pos := sort.SearchInts(F, nums[(n - 1) / 2])
    if F[pos] <= nums[n / 2] {
        return int64(calc(F[pos]))
    }

    return int64(min(calc(F[pos]), calc(F[pos - 1])))
}

func abs(x int) int { if x < 0 { return -x }; return x }