use std::cmp::max;
impl Solution {
    const MOD: i64 = 1_000_000_007;
    pub fn max_sum(a: Vec<i32>, b: Vec<i32>) -> i32 {
        let (mut i, mut j) = (0, 0);
        let (mut x, mut y) = (0i64, 0i64);
        let m = a.len();
        let n = b.len();
        while i < m || j < n {
            if i < m && j < n {
                if a[i] < b[j] {
                    x += a[i] as i64;
                    i += 1;
                } else if a[i] > b[j] {
                    y += b[j] as i64;
                    j += 1;
                } else {
                    let t = max(x, y) + a[i] as i64;
                    x = t;
                    y = t;
                    i += 1;
                    j += 1;
                }
            } else if i < m {
                x += a[i] as i64;
                i += 1;
            } else {
                y += b[j] as i64;
                j += 1;
            }
        }
        (max(x, y) % Self::MOD) as i32
    }

}