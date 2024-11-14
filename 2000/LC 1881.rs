impl Solution {
    pub fn max_value(s: String, x: i32) -> String {
        let n = s.len();
        let s = &s;
        let mut i = 0;

        if s.chars().next() == Some('-') {
            i += 1;
            while i < n {
                if s[i..i+1].parse::<i32>().unwrap() > x {
                    break;
                }
                i += 1;
            }
        } else {
            while i < n {
                if s[i..i+1].parse::<i32>().unwrap() < x {
                    break;
                }
                i += 1;
            }
        }

        let mut res = s[0..i].to_string();
        res.push_str(&x.to_string());
        res.push_str(&s[i..]);

        res
    }
}