impl Solution {
    pub fn can_form_array(arr: Vec<i32>, pieces: Vec<Vec<i32>>) -> bool {
        let mut pos = vec![-1; 101];
        arr.iter().enumerate().for_each(|(i, v)| pos[*v as usize] = i as i32);
        for v in pieces {
            if v.len() == 1 {
                if pos[v[0] as usize] == -1 {
                    return false;
                }
            } else {
                for i in 1..v.len() {
                    if pos[v[i] as usize] != pos[v[i-1] as usize] + 1 {
                        return false;
                    }
                }
            }
        }

        true
    }
}