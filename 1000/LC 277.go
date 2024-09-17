/* The knows API is defined for you.
       knows(a: i32, b: i32)->bool;
    to call it use self.knows(a,b)
*/

impl Solution {
    pub fn find_celebrity(&self, n: i32) -> i32 {
        let mut ans = 0;
        for i in 0..n {
            if self.knows(ans, i) {
                ans = i
            }
        }

        for i in 0..n {
            if i == ans {
                continue
            } else {
                if self.knows(ans, i) || !self.knows(i, ans) {
                    return -1;
                }
            }
        }

        ans as i32
    }
}