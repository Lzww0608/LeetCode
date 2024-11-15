impl Solution {
    pub fn count_good_numbers(n: i64) -> i32 {
        const MOD: i64 = 1_000_000_007;
        let quickPow = |mut a: i64, mut r: i64| -> i64 {
            let mut res = 1;
            while r > 0 {
                if (r & 1) == 1 {
                    res = (res * a) % MOD;
                }
                r >>= 1;
                a = (a * a) % MOD;
            }
            res
        };

        let a = quickPow(4, n / 2);
        let b = quickPow(5, (n + 1) / 2);
        ((a * b) % MOD) as _
    }
}