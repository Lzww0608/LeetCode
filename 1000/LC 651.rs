impl Solution {
    pub fn max_a(n: i32) -> i32 {
        let n = n as usize;
        let mut f = vec![0; n + 1];
        f[1] = 1;
        for i in 2..=n {
            f[i] = f[i-1] + 1;
            for j in 0..(i-2) {
                f[i] = f[i].max(f[j] * (i - (j + 2) + 1));
            }
        }

        f[n] as _
    }
}