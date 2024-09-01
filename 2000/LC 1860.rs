impl Solution {
    pub fn mem_leak(mut memory1: i32, mut memory2: i32) -> Vec<i32> {
        let mut t: i32 = 0;
        while memory1 >= t || memory2 >= t {
            if memory1 >= memory2 {
                memory1 -= t;
            } else {
                memory2 -= t;
            }
            t += 1;
        }

        vec![t, memory1, memory2]    
    }
}