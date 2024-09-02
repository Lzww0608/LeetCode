impl Solution {
    pub fn appeal_sum(s: String) -> i64 {
        let mut ans: i64 = 0;
        let mut sum: i64 = 0;
        let mut pos = vec![-1; 26];

        for (i, c) in s.chars().enumerate() {
            let x: u8 = c as u8 - b'a';
            sum += i as i64 - pos[x as usize];
            ans += sum;
            pos[x as usize] = i as i64;
        }

        ans
    }
}