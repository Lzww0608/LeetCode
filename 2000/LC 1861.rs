impl Solution {
    pub fn rotate_the_box(box_: Vec<Vec<char>>) -> Vec<Vec<char>> {
        let m = box_.len();
        let n = box_[0].len();
        let mut ans = vec![vec!['.'; m]; n];

        for i in 0..m {
            let mut cnt = 0;
            for j in 0..n {
                if box_[i][j] == '#' {
                    cnt += 1;
                } else if box_[i][j] == '*' {
                    ans[j][m - 1 - i] = '*';
                    let mut k = j as i32 - 1;
                    while cnt > 0 {
                        cnt -= 1;
                        ans[k as usize][m - 1 - i] = '#';
                        k -= 1;
                    }
                }
            }

            let mut k = n as i32 - 1;
            while cnt > 0 {
                cnt -= 1;
                ans[k as usize][m - 1 - i] = '#';
                k -= 1;
            }
        }

        ans
    }
}