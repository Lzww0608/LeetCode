const N: usize = 26;
impl Solution {
    pub fn can_permute_palindrome(s: String) -> bool {
        let mut cnt = vec![0; N];
        let s = s.as_bytes();
        for i in 0..s.len() {
            cnt[(s[i]-b'a') as usize] += 1;
        }

        let mut sum = 0;
        for &x in &cnt {
            if x & 1 == 1 {
                sum += 1;
            }
        }

        sum <= 1
    }
}