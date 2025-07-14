
import (
	"unicode"
)
func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }

    vowel := map[rune]bool {
        'a': true,
        'e': true, 
        'i': true, 
        'o': true,
        'u': true,
        'A': true,
        'E': true, 
        'I': true, 
        'O': true,
        'U': true,
    }

    mask := 0
    for _, c := range word {
        if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
            return false
        }

        if unicode.IsDigit(c) {
            continue
        }

        if vowel[c] {
            mask |= 1
        } else {
            mask |= 2
        }
    }

    return mask == 3
}