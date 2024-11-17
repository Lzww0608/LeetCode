impl Solution {
    pub fn multiply(a: Vec<Vec<i32>>, b: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let m = a.len();
        let k = a[0].len();
        let n = b[0].len();

        let mut ans = vec![vec![0; n]; m];
        
        for i in 0..m {
            for j in 0..n {
                for p in 0..k {
                    ans[i][j] += a[i][p] * b[p][j];
                }
            }
        }

        ans
    }
}