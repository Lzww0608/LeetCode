func superpalindromesInRange(left string, right string) (ans int) {
	l, _ := strconv.Atoi(left)
	r, _ := strconv.Atoi(right)

	limit := int(math.Sqrt(float64(r)))
	seed := 1
	num := 0

	for num <= limit {
		num = constructPalindrome(seed, false)
		if check(num*num, l, r) {
			ans++
		}
		num = constructPalindrome(seed, true)
		if check(num*num, l, r) {
			ans++
		}
		seed++
	}
	return ans
}

func constructPalindrome(seed int, odd bool) int {
	ans := seed
    if odd {
        seed /= 10
    }
	for seed != 0 {
		ans = ans*10 + seed%10
		seed /= 10
	}
	return ans
}

func check(x, l, r int) bool {
	return x >= l && x <= r && isPalindrome(x)
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
    n := len(s)
    for i := 0; i + i < n; i++ {
        if s[i] != s[n-i-1] {
            return false
        }
    }

    return true
}