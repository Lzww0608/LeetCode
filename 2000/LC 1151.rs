impl Solution {
    pub fn min_swaps(data: Vec<i32>) -> i32 {
        let mut cnt = 0;
        for &x in &data {
            if x == 1 {
                cnt += 1;
            }
        }

        let mut cur = 0;
        let mut ans = cnt;
        let n = data.len();
        for i in 0..n {
            if i >= cnt as usize && data[i- cnt as usize] == 1 {
                cur -= 1;
            }
            if data[i] == 1 {
                cur += 1;
            }

            ans = ans.min(cnt - cur);
        }

        ans as i32
    }
}