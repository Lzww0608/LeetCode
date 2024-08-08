impl Solution {
    pub fn min_score_triangulation(values: Vec<i32>) -> i32 {
        let n: usize = values.len();
        let mut f = vec![vec![0; n]; n];
        
        for i in (0..=n - 3).rev() {
            for j in i+2..n {
                let mut t: i32 = std::i32::MAX;
                for k in i+1..j {
                    t = t.min(values[i] * values[j] * values[k] + f[i][k] + f[k][j])
                }
                f[i][j] = t
            }
        }

        f[0][n-1]
    }
}