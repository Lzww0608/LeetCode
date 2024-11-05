
impl Solution {
    pub fn max_consecutive(bottom: i32, top: i32, mut special: Vec<i32>) -> i32 {
        special.sort_unstable();
        let mut pre = bottom - 1;
        let mut ans = 0;
        for &x in &special {
            ans = ans.max(x - pre - 1);
            pre = x;
        }
        
        ans.max(top - pre)
    }
}