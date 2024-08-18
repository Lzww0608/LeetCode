func cmp(n *big.Int, digit int, base *big.Int) int {
	cur := big.NewInt(0)
	one := big.NewInt(1)
	for i := 0; i < digit; i++ {
		cur.Mul(cur, base)
		cur.Add(cur, one)
		if cur.Cmp(n) > 0 {
			return 1
		}
	}
	if cur.Cmp(n) < 0 {
		return -1
	}
	return 0
}

func smallestGoodBase(s string) string {
	n := new(big.Int)
	n, _ = n.SetString(s, 10)

	for i := 64; i >= 2; i-- {
		left := big.NewInt(2)
		right := new(big.Int).Sub(n, big.NewInt(1))
		mid := big.NewInt(0)
		
		for left.Cmp(right) <= 0 {
			mid.Add(left, right)
			mid.Div(mid, big.NewInt(2))

			chk := cmp(n, i, mid)
			if chk == 0 {
				return mid.String()
			} else if chk > 0 {
				right.Sub(mid, big.NewInt(1))
			} else {
				left.Add(mid, big.NewInt(1))
			}
		}
	}
	return new(big.Int).Sub(n, big.NewInt(1)).String()
}