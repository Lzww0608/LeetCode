impl Solution {
    pub fn min_flips(s: String) -> i32 {
        let n = s.len();
        let target = ['0', '1'];
        let mut cnt = 0;
        let s: Vec<char> = s.chars().collect();

        for i in 0..n {
            if s[i] != target[i & 1] {
                cnt += 1;
            }
        }

        let mut ans = cnt.min(n - cnt);
        for i in 0..n {
            if s[i] != target[i & 1] {
                cnt -= 1;
            }
            if s[i] != target[(i + n) & 1] {
                cnt += 1;
            }
            ans = ans.min(cnt.min(n - cnt));
        }

        ans as i32
    }
}