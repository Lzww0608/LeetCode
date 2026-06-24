const MOD int = 1_000_000_007

func zigZagArrays(n int, l int, r int) int {
	k := r - l + 1
	f := make([]int, k)
	for i := range f {
		f[i] = 1
	}

	a := make([][]int, k)
	b := make([][]int, k)
	for i := range a {
		a[i] = make([]int, k)
		b[i] = make([]int, k)
		for j := range k {
			if j < i {
				a[i][j] = 1
			} else if j > i {
				b[i][j] = 1
			}
		}
	}

	p := mulMat(b, a)
	tmp := powMat(p, (n-1)/2)
	f = mulVec(tmp, f)
	if (n-1)&1 == 1 {
		f = mulVec(a, f)
	}

	ans := 0
	for _, v := range f {
		ans = (ans + v*2) % MOD
	}
	return ans
}

func mulVec(a [][]int, b []int) []int {
	res := make([]int, len(a))
	for i := range res {
		for j := range a[0] {
			res[i] = (res[i] + a[i][j]*b[j]) % MOD
		}
	}
	return res
}

func mulMat(a, b [][]int) [][]int {
	res := make([][]int, len(a))
	for i := range res {
		res[i] = make([]int, len(b[0]))
	}
	for i := range a {
		for j := range b[0] {
			for k := range a[0] {
				res[i][j] = (res[i][j] + a[i][k]*b[k][j]) % MOD
			}
		}
	}
	return res
}

func powMat(a [][]int, n int) [][]int {
	res := make([][]int, len(a))
	for i := range res {
		res[i] = make([]int, len(a[0]))
	}
	for i := range res {
		res[i][i] = 1
	}
	for n > 0 {
		if n%2 == 1 {
			res = mulMat(res, a)
		}
		a = mulMat(a, a)
		n >>= 1
	}
	return res
}