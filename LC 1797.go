type AuthenticationManager struct {

}

var liveTime int 
var m map[string]int

func Constructor(timeToLive int) AuthenticationManager {
    liveTime = timeToLive
    m = make(map[string]int)
    return AuthenticationManager{}
}


func (this *AuthenticationManager) Generate(tokenId string, currentTime int)  {
    m[tokenId] = currentTime + liveTime

}


func (this *AuthenticationManager) Renew(tokenId string, currentTime int)  {
    if m[tokenId] > currentTime {
        m[tokenId] = currentTime + liveTime
    }
}


func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
    cnt := 0
    for k, v := range m {
        if v <= currentTime {
           delete(m, k) 
        } else {
            cnt++
        }
    }
    return cnt
}


/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */