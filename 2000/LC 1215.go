func countSteppingNumbers(low int, high int) []int {
	ans := []int{0}
	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} 

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr >= low && curr <= high {
			ans = append(ans, curr)
		}

		if curr > high { 
			continue
		}

		lastDigit := curr % 10
		if lastDigit > 0 {
			nextNum := curr*10 + (lastDigit - 1)
			q = append(q, nextNum)
		}
		if lastDigit < 9 {
			nextNum := curr*10 + (lastDigit + 1)
			q = append(q, nextNum)
		}
	}

	//sort.Ints(ans)
    if low > 0 {
        ans = ans[1:]
    }
	return ans
}