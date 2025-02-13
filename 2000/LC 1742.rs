use std::collections::HashMap;
impl Solution {
    pub fn count_balls(low_limit: i32, high_limit: i32) -> i32 {
        let mut cnt: HashMap<i32, i32> = HashMap::new();
        
        let solve = |mut x: i32| -> i32 {
            let mut res = 0;
            while x > 0 {
                res += x % 10;
                x /= 10;
            }
            res
        };

        let mut ans = 0;
        for x in low_limit..=high_limit {
            let mut t = solve(x);
            *cnt.entry(t).or_insert(0) += 1;
            ans = ans.max(*cnt.get(&t).unwrap());
        }

        ans 
    }
}