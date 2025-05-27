impl Solution {
    pub fn encode(num: i32) -> String {
        if num == 0 {
            return "".to_string();
        }
        let s = format!("{:b}", num + 1);
        s[1..].to_string()
    }
}