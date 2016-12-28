package cross

func IsOk(value int, valid []int)(isok bool){
    isok=false
    for i:=range valid { if value==valid[i] { isok=true ; break } }
    return
}
