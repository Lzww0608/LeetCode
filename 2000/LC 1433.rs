impl Solution {
    pub fn check_if_can_break(s: String, t: String) -> bool {
        let n = s.len();
        let mut cnt_s = vec![0; 26];
        let mut cnt_t = vec![0; 26];
         for i in 0..n {
            let a = (s.as_bytes()[i] - b'a') as usize;
            let b = (t.as_bytes()[i] - b'a') as usize;
            cnt_s[a] += 1;
            cnt_t[b] += 1;
        }
        
        let mut f1 = true;
        let mut f2 = true;
        let mut sumS = 0;
        let mut sumT = 0;
        for i in 0..26 {
            sumS += cnt_s[i];
            sumT += cnt_t[i];
            if sumS > sumT {
                f2 = false;
            } else if sumT > sumS {
                f1 = false;
            }

            if !f1 && !f2 {
                return false;
            }
        }

        true
    }
}