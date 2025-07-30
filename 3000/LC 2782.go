/**
 * Definition for a category handler.
 * type CategoryHandler interface {
 *  HaveSameCategory(int, int) bool
 * }
 */
func numberOfCategories(n int, categoryHandler CategoryHandler) (ans int) {
    for i := 0; i < n; i++ {
        f := true 
        for j := i - 1; j >= 0; j-- {
            if categoryHandler.HaveSameCategory(i, j) {
                f = false
                break
            }
        } 
        if f {
            ans++
        }
    }

    return
}