func game(guess []int, answer []int) (ans int) {
    for i := range guess {
       if guess[i] == answer[i] {
            ans++
       } 
    }

    return
}