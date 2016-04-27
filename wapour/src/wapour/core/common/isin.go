package common
func IsIn(x string,xs []string)(bool) {
    for i := range xs {
        y:= xs[i]
        if y == x {return true}
    }
    return false
}
