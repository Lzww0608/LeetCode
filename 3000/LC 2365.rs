use std::collections::HashMap;

impl Solution {
    pub fn task_scheduler_ii(tasks: Vec<i32>, space: i32) -> i64 {
        let mut ans = 0;
        let mut last: HashMap<i32, i64> = HashMap::new();

        for &t in &tasks {
            ans += 1;
            if let Some(&v) = last.get(&t) {
                ans = ans.max(v + space as i64 + 1);
            }
            last.insert(t, ans);
        }

        ans
    }
}