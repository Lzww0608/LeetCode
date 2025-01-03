const N: usize = 26;
impl Solution {
    pub fn sort_string(s: String) -> String {
        let mut cnt = vec![0; N];
        for c in s.chars() {
            cnt[(c as usize) - ('a' as usize)] += 1;
        }

        let mut ans = String::new();
        let n = s.len();

        while ans.len() < n {
            let mut a = Vec::new();
            let mut b = Vec::new();

            for i in 0..N {
                if cnt[i] > 0 {
                    cnt[i] -= 1;
                    a.push((i as u8 + 'a' as u8) as char);
                }
            }

            for i in (0..N).rev() {
                if cnt[i] > 0 {
                    cnt[i] -= 1;
                    b.push((i as u8 + 'a' as u8) as char);
                }
            }

            ans.push_str(&a.iter().collect::<String>());
            ans.push_str(&b.iter().collect::<String>());
        }

        ans
    }
}