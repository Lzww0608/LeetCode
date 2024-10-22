const N: i32 = 24;
impl Solution {
    pub fn count_complete_day_pairs(hours: Vec<i32>) -> i64 {
        let mut ans = 0;
        let mut cnt = vec![0; N as usize];

        for &x in &hours {
            let x = x % N;
            ans += cnt[((N-x)%N) as usize] as i64;
            cnt[x as usize] += 1;
        }

        ans
    }
}