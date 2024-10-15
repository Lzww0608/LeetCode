impl Solution {
    pub fn smallest_common_element(mat: Vec<Vec<i32>>) -> i32 {
        let m = mat.len();
        let n = mat[0].len();
        let mut p = vec![0; m];

        for j in 0..n {
            let x = mat[0][j];
            let mut f = true;
            for j in 1..m {
                while p[j] < n && mat[j][p[j]] < x {
                    p[j] += 1;
                }

                if p[j] >= n || mat[j][p[j]] > x {
                    f = false;
                }
            }
            if f {
                return x;
            }
        }

        -1
    }
}