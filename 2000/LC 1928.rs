use std::cmp::min;
impl Solution {
    pub fn min_cost(max_time: i32, edges: Vec<Vec<i32>>, passing_fees: Vec<i32>) -> i32 {
        let n = passing_fees.len();
        let max_time = max_time as usize;
        let mut f = vec![vec![i32::MAX / 2; n]; max_time + 1];
        f[0][0] = passing_fees[0];
        for i in 1..=max_time {
            for e in &edges {
                let u = e[0] as usize;
                let v = e[1] as usize;
                let t = e[2] as usize;
                if i >= t {
                    f[i][u] = min(f[i][u], f[i-t][v] + passing_fees[u]);
                    f[i][v] = min(f[i][v], f[i-t][u] + passing_fees[v]);
                }
            }
        }

        let mut ans = i32::MAX;
        for i in 1..=max_time {
            ans = min(ans, f[i][n - 1]);
        }
        
        if ans >= i32::MAX / 2 {
            -1
        } else {
            ans
        }
    }
}
