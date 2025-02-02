impl Solution {
    pub fn reformat(s: String) -> String {
        let n = s.len();
        let (mut digits, mut letters): (Vec<_>, Vec<_>) = s.chars().partition(|c| c.is_ascii_digit());

        if digits.len() < letters.len() {
            std::mem::swap(&mut digits, &mut letters);
        }

        if digits.len() as i32 - letters.len() as i32 > 1 {
            "".to_string()
        } else {
            let mut ans = Vec::with_capacity(n);
            for (&a, &b) in digits.iter().zip(&letters) {
                ans.push(a);
                ans.push(b);
            }
            if digits.len() > letters.len() {
                ans.push(*digits.last().unwrap() as char);
            }

            ans.into_iter().collect()
        }
    }
}