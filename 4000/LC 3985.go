func ManacherNums(nums []int) []int {
	n := len(nums)
	m := 2*n + 1
	p := make([]int, m)

	center, right := 0, 0

	for i := 0; i < m; i++ {
		mirror := 2*center - i

		if i < right && mirror >= 0 {
			p[i] = min(p[mirror], right-i)
		}

		for i-p[i]-1 >= 0 &&
			i+p[i]+1 < m &&
			equalVirtual(nums, i-p[i]-1, i+p[i]+1) {
			p[i]++
		}

		if i+p[i] > right {
			center = i
			right = i + p[i]
		}
	}

	return p
}

func equalVirtual(nums []int, a, b int) bool {
	aIsSep := a%2 == 0
	bIsSep := b%2 == 0

	if aIsSep || bIsSep {
		return aIsSep && bIsSep
	}

	return nums[a/2] == nums[b/2]
}

func getSum(nums []int) int64 {
    n := len(nums)
    p := ManacherNums(nums)
    sum := make([]int, n + 1)
    for i, x := range nums {
        sum[i + 1] = sum[i] + x
    }

    ans := 0
    for i, x := range p {
        if x == 0 {
            continue
        }

        l, r := (i - x) / 2, (i + x) / 2 - 1
        ans = max(ans, sum[r + 1] - sum[l])
    }

    return int64(ans)
}