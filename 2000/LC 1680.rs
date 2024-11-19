impl Solution {
    pub fn concatenated_binary(n: i32) -> i32 {
        let MOD: i64 = 1_000_000_007;
        let mut ans: i64 = 0;
        let mut shift = 0;
        for i in 1..=n {
            if i & (i - 1) == 0 {
                shift += 1;
            }
            ans = ((ans << shift) + i as i64) % MOD;
        }

        ans as _
    }
}