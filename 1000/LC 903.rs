const MOD: i32 = 1_000_000_007;

impl Solution {
    pub fn num_perms_di_sequence(s: String) -> i32 {
        let n = s.len();
        let mut f = vec![vec![0; n + 1]; n + 1];
        f[0][0] = 1;

        for i in 1..=n {
            if &s[i-1..i] == "D" {
                f[i][i] = 0;
                for j in (0..i).rev() {
                    f[i][j] = (f[i][j+1] + f[i-1][j]) % MOD;
                }
            } else {
                f[i][0] = 0;
                for j in 1..=i {
                    f[i][j] = (f[i][j-1] + f[i-1][j-1]) % MOD;
                }
            }
        }

        let mut ans = 0;
        for j in 0..=n {
            ans = (ans + f[n][j]) % MOD;
        }

        ans
    }
}