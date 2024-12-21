impl Solution {
    pub fn array_strings_are_equal(word1: Vec<String>, word2: Vec<String>) -> bool {
        let a: String = word1.concat();
        let b: String = word2.concat();

        a == b
    }
}