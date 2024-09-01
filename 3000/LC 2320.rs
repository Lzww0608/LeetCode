impl Solution {
    const MOD: i64 = 1_000_000_007;
    pub fn count_house_placements(n: i32) -> i32 {
        let n = n as usize;
        let mut f = vec![0; n + 1];
        f[0] = 1;
        f[1] = 2;

        for i in 2..n+1 {
            f[i] = f[i-1] + f[i-2];
            f[i] %= Self::MOD;
        }

        ((f[n] * f[n]) % Self::MOD) as i32
    }
}