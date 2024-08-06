impl Solution {
    pub fn max_side_length(mat: Vec<Vec<i32>>, threshold: i32) -> i32 {
        let mut ans: i32 = 0;
        let m = mat.len();
        let n = mat[0].len();

        let mut sum = vec![vec![0; n + 1]; m + 1];

        for i in 0..m {
            for j in 0..n {
                sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + mat[i][j];
            }
        }

        for i in 1..=m {
            for j in 1..=n {
                for k in (ans + 1) as usize..=i {
                    if k <= j {
                        let cur = sum[i][j] - sum[i-k][j] - sum[i][j-k] + sum[i-k][j-k];
                        if cur <= threshold {
                            ans = ans.max(k as i32);
                        } else {
                            break;
                        }
                    }
                }
            }
        }

        ans
    }
}