
impl Solution {
    pub fn add_spaces(s: String, spaces: Vec<i32>) -> String {
        let s = s.as_str();
        let mut builder = String::new();
        let mut pre: usize = 0;

        for &x in &spaces {
            builder.push_str(&s[pre..x as usize]);
            builder.push(' ');
            pre = x as usize;
        }

        builder.push_str(&s[pre..]);

        builder
    }
}