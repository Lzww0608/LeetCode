func rotateString(s string, goal string) bool {
    t := s + s 
    return len(s) == len(goal) && strings.Contains(t, goal)
}