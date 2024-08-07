impl Solution {
    pub fn maximum_rows(matrix: Vec<Vec<i32>>, num_select: i32) -> i32 {
        let m = matrix.len();
        let n = matrix[0].len();
        let mut ans: i32 = 0;

        let mut mask = vec![0; m];
        for i in 0..m {
            for j in 0..n {
                mask[i] |= (matrix[i][j] << j);
            }
        }

        let mut s = (1 << num_select) - 1;
        while s < (1 << n) {
            let mut t: i32 = 0;
            for &k in &mask {
                if (k & s) == k {
                    t += 1;
                }
            }

            ans = ans.max(t);

            let lb = s & (-s);
            let x = s + lb;
            s = ((s ^ x) / lb >> 2) | x;
        }

        ans
    }
}