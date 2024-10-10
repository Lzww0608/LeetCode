impl Solution {
    pub fn probability_of_heads(prob: Vec<f64>, target: i32) -> f64 {
        let n = prob.len();
        let target = target as usize;
        let mut f = vec![vec![0f64; target+1]; n+1];
        f[0][0] = 1.0;
        
        for i in 1..=n {
            f[i][0] = f[i-1][0] * (1.0 - prob[i-1]);
            for j in 1..=target {
                f[i][j] = f[i-1][j-1] * prob[i-1] + f[i-1][j] * (1.0 - prob[i-1]);
            }
        }

        f[n][target]
    }
}