const MOD: i64 = 1_000_000_007;
impl Solution {
    pub fn rearrange_sticks(n: i32, k: i32) -> i32 {
        let n = n as usize;
        let k = k as usize;
        let mut f = vec![vec![0; n + 1]; n + 1];
        f[0][0] = 1;

        for i in 1..=n {
            for j in 1..=i {
                f[i][j] = (f[i-1][j-1] + (i as i64 - 1) * f[i-1][j]) % MOD;
            }
        }

        f[n][k] as i32
    }
}