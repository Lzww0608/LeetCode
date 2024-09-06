use std::cmp::min;
impl Solution {
    pub fn min_cost(costs: Vec<Vec<i32>>) -> i32 {
        let (m, n) = (costs.len(), costs[0].len());
        let mut f = vec![vec![i32::MAX / 2; n]; m];
        
        for i in 0..n {
            f[0][i] = costs[0][i];
        }

        for i in 1..m {
            for j in 0..n {
                for k in 0..n {
                    if k != j {
                        f[i][j] = min(f[i][j], f[i-1][k] + costs[i][j]);
                    }
                }
            }
        }

        *f[m-1].iter().min().unwrap()
    }
}