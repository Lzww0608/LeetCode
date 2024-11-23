impl Solution {
    pub fn matrix_sum_queries(n: i32, queries: Vec<Vec<i32>>) -> i64 {
        let m = queries.len();
        let mut ans: i64 = 0;
        let (mut cnt_row, mut cnt_col) = (0, 0);
        let mut row = vec![false; n as usize];
        let mut col = vec![false; n as usize];

        for i in (0..m).rev() {
            let q = &queries[i];
            let k = q[1] as usize;
            if q[0] == 0 {
                if !row[k] {
                    ans += q[2] as i64 * (n - cnt_col) as i64;
                    cnt_row += 1;
                    row[k] = true;
                }
            } else {
                if !col[k] {
                    ans += q[2] as i64 * (n - cnt_row) as i64;
                    cnt_col += 1;
                    col[k] = true;
                }
            }
        }

        ans
    }
}