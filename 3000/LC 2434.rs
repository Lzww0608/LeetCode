impl Solution {
    pub fn robot_with_string(s: String) -> String {
        let mut cnt = vec![0; 26];
        let mut st: Vec<u8> = Vec::new();
        let mut ans: Vec<u8> = Vec::with_capacity(s.len());
        let s = s.as_bytes();
        let n = s.len();
        for i in 0..n {
            let x = (s[i] & 31) as usize - 1;
            cnt[x] += 1;
        }

        let mut mn = 0;
        while cnt[mn] == 0 {
            mn += 1;
        }
        for i in 0..n {
            st.push(s[i]);
            let x = (s[i] & 31) as usize - 1;
            cnt[x] -= 1;
            while mn < 26 && cnt[mn] == 0 {
                mn += 1;
            }

            while let Some(&last) = st.last() {
                if last <= b'a' + mn as u8 {
                    ans.push(last);
                    st.pop();
                } else {
                    break;
                }
            }
        }

        String::from_utf8(ans).unwrap()
    }
}