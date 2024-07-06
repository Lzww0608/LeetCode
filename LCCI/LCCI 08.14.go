func countEval(s string, result int) int {
    memo := make(map[string]int)
	return countWaysHelper(s, result, memo)
}

func countWaysHelper(s string, result int, memo map[string]int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		if int(s[0]-'0') == result {
			return 1
		}
		return 0
	}

	key := fmt.Sprintf("%s_%d", s, result)
	if val, found := memo[key]; found {
		return val
	}

	ways := 0
	for i := 1; i < len(s); i += 2 {
		left := s[:i]
		right := s[i+1:]
		operator := s[i]

		leftTrue := countWaysHelper(left, 1, memo)
		leftFalse := countWaysHelper(left, 0, memo)
		rightTrue := countWaysHelper(right, 1, memo)
		rightFalse := countWaysHelper(right, 0, memo)

		totalWays := (leftTrue + leftFalse) * (rightTrue + rightFalse)

		trueWays := 0
		switch operator {
		case '&':
			trueWays = leftTrue * rightTrue
		case '|':
			trueWays = leftTrue*rightTrue + leftTrue*rightFalse + leftFalse*rightTrue
		case '^':
			trueWays = leftTrue*rightFalse + leftFalse*rightTrue
		}

		if result == 1 {
			ways += trueWays
		} else {
			ways += totalWays - trueWays
		}
	}

	memo[key] = ways
	return ways
}