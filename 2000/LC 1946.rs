impl Solution {
    pub fn maximum_number(num: String, change: Vec<i32>) -> String {
        let n = num.len();
        let mut ans: Vec<char> = num.chars().collect();

        let mut f = false;
        for i in 0..n {
            let x: i32 = ans[i] as i32 & 15;
            if x < change[x as usize] || (f && x == change[x as usize]) {
                f = true;
                ans[i] = (b'0' + change[x as usize] as u8) as char;
            } else if f {
                break;
            }
        }

        ans.into_iter()
            .collect()
    }
}