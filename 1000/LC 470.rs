/** 
 * The rand7() API is already defined for you.
 * @return a random integer in the range 1 to 7
 * fn rand7() -> i32;
 */

impl Solution {
    pub fn rand10() -> i32 {
        while true {
            let mut ans = (rand7() - 1) * 7 + rand7();
            if ans <= 40 {
                return ans % 10 + 1
            }
        }

        0
    }
}